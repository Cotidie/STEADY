from datetime import datetime
from flask import Blueprint, jsonify, request

api_miniter = Blueprint('miniter', import_name=__name__)

SIMPLE_DB = {
    'users': {},     # id가 key
    'articles': {},  # id가 key, value는 {datetime, content}
    'follows': {},   # id가 key, value는 {ids} (set)
}


@api_miniter.route('/ping', methods=['GET', 'POST'])
def health_check():
    return "pong"


# 1. 회원가입하기
# * json으로 받기, 보내기
@api_miniter.route('/sign-up', methods=['POST'])
def sign_up():
    """
    회원가입 API. id, name, email, password, profile로 구성된 json이 필요하다.
    :return: json { 'status': true|false }
    """
    req_data = request.get_json()
    res_data = {'status': False}
    if (not req_data):
        return jsonify(res_data), 400

    new_user = {
        'name'      : req_data['name'],
        'password'  : req_data['password'],
        'profile'   : req_data['profile']
    }
    SIMPLE_DB['users'][int(req_data['id'])] = new_user

    print(new_user)

    return jsonify(new_user)


# 2. 300자 제한 글 올리기
# * Response에 status code 명시하기
@api_miniter.route('/post', methods=['POST'])
def post_article():
    """
    - 사용자는 300자를 초과하지 않는 글을 올릴 수 있다.
    - 300자를 초과한다면 400 Bad Request 응답을 내보내야 한다.
    - 300자 이내의 글을 전송하면 글을 저장하고, 유저에 따라 타임라인을 볼 수 있어야 한다.
    - id, content로 구성된 json이 필요하다.
    """
    payload = request.get_json()
    res_data = {'status': False}
    if not payload:
        return jsonify(res_data), 400

    user_id = int(payload['id'])
    # id 존재여부 검사
    if not SIMPLE_DB['users'].get(user_id):
        return jsonify(res_data), 400

    content = str(payload['content'])
    # 300자 초과 검사
    if len(content) > 300:
        return jsonify(res_data), 400

    new_article = {
        'datetime': datetime.now(),
        'content': content
    }

    timeline = SIMPLE_DB['articles'].get(user_id)
    if not timeline:
        SIMPLE_DB['articles'][user_id] = [new_article]
    else:
        SIMPLE_DB['articles'][user_id].append(new_article)

    return user_id, 200


# 3. 팔로우 기능
@api_miniter.route('/follow', methods=['POST'])
def follow():
    """
    {id, follow}로 이루어진 JSON
    """
    payload     = request.get_json()
    user_id     = int(payload['id'])
    to_follow   = int(payload['follow'])

    res_data = {'status': False}

    if (user_id not in SIMPLE_DB['users']) or \
        (to_follow not in SIMPLE_DB['users']):
        return jsonify(res_data), 400

    SIMPLE_DB['follows'].setdefault(user_id, set()).add(to_follow)

    return user_id, 200


# 4. 언팔로우
@api_miniter.route('/unfollow', methods=['POST'])
def unfollow():
    """
    {id, unfollow로 이루어진 JSON
    """
    payload     = request.get_json()
    user_id     = payload['id']
    to_unfollow = payload['unfollow']

    res_data = {'status': False}

    if (not user_id) or (not to_unfollow):
        return jsonify(res_data), 400

    if (user_id not in SIMPLE_DB['users']) or \
        (to_unfollow not in SIMPLE_DB['users']):
        return jsonify(res_data), 400

    SIMPLE_DB['follows'].setdefault(user_id, set()).discard(to_unfollow)

    return user_id, 200


# 5. 유저 정보 얻기
@api_miniter.route('/user/<int:user_id>', methods=['GET'])
def get_user(user_id: int):
    user = SIMPLE_DB['users'].get(user_id)
    if not user:
        return "No such user", 400

    return jsonify(user), 200
