from datetime import datetime, timedelta
from functools import wraps
import bcrypt
import jwt
from flask import Blueprint, jsonify, request, current_app as app, Response
from pymongo.errors import DuplicateKeyError
from bcrypt import hashpw

api_miniter = Blueprint('miniter', import_name=__name__)
SIMPLE_DB = {}

# 인증 데코레이터
def login_required(f):
    @wraps(f)
    def decorated_func(*args, **kwargs):
        access_token = request.headers.get('Authorization')
        if not access_token:
            # 토큰이 없는 경우
            return Response(status=401)

        # 토큰 복호화 시도
        try:
            payload = jwt.decode(access_token,
                                 app.config.get('JWT_SECRETKEY'),
                                 'HS256')
        except jwt.InvalidTokenError:
            payload = None

        # 유효하지 않은 토큰
        if not payload:
            return Response(status=401) # 401: 인증 오류

        return f(*args, **kwargs)
    return decorated_func


@api_miniter.route('/ping', methods=['GET', 'POST'])
def health_check():
    return "pong"


# 회원가입하기
@api_miniter.route('/sign-up', methods=['POST'])
def sign_up():
    """
    회원가입 API. id, name, email, password, profile로 구성된 json이 필요하다.
    * 비밀번호는 bcrypt로 암호화된다.
    * email 필드는 unique index로 설정되어 중복이 불가능하다.
    :return: json { 'status': true|false }
    """
    req_data = request.get_json()
    res_data = {
        'status': False,
        'message': ""
    }
    if (not req_data):
        return jsonify(res_data), 400

    new_user = {
        'name'      : req_data['name'],
        'password'  : req_data['password'],
        'email'     : req_data['email'],
        'profile'   : req_data['profile']
    }

    try:
        app.db.add_new_user(new_user)
    except DuplicateKeyError:
        res_data['message'] = "이메일이 중복됩니다."
        return jsonify(res_data), 400

    return jsonify(new_user), 200

# 로그인하기
@api_miniter.route('/login', methods=['POST'])
def login():
    """
    로그인 API. 로그인에 성공하면 access token을 발급한다.
    - 들어오는 JSON은 'email', 'password'로 구성된다.
    - pyJWT로 access token을 발급한다.
    """
    # 데이터 받기
    credential = request.get_json()
    user_id     = str(credential.get('email'))
    password    = str(credential.get('password'))

    # 결과물
    res_data = {
        'status': False,
        'message': ""
    }

    # 오류 검사
    if not user_id or not password:
        res_data['message'] = "데이터가 올바르지 않습니다."
        return jsonify(res_data), 400

    # 유저 데이터
    user = app.db.get_user(user_id)
    cred = app.db.get_cred(user_id)
    if not user or not cred:
        res_data['message'] = "유저가 존재하지 않습니다."
        return jsonify(res_data), 400

    # 패스워드 검증
    hashed = hashpw(password.encode(encoding='utf-8'), cred['salt'])
    if hashed != cred['password']:
        res_data['message'] = "패스워드가 올바르지 않습니다."
        return jsonify(res_data), 400

    # 토큰 발급
    payload = {
        'email': user_id,
        'exp': datetime.now() + timedelta(days=365)
    }
    key = app.config.get('JWT_SECRETKEY')
    token = jwt.encode(payload, key)

    res_data['status'] = True
    res_data['access_token'] = token

    return jsonify(res_data), 200


# 300자 제한 글 올리기
# * Response에 status code 명시하기
@api_miniter.route('/article', methods=['POST'])
@login_required
def create_article():
    """
    - 사용자는 300자를 초과하지 않는 글을 올릴 수 있다.
    - 300자를 초과한다면 400 Bad Request 응답을 내보내야 한다.
    - 300자 이내의 글을 전송하면 글을 저장하고, 유저에 따라 타임라인을 볼 수 있어야 한다.
    - email, content로 구성된 json이 필요하다.
    """
    payload = request.get_json()
    res_data = {'status': False}
    if not payload:
        return jsonify(res_data), 400

    email = str(payload['email']); print(email)
    content = str(payload['content'])
    user = app.db.get_user(email)
    if not user:
        res_data['message'] = '존재하지 않는 유저입니다.'
        return jsonify(res_data), 400
    print('checkpoint2')
    # 300자 초과 검사
    if len(content) > 300:
        res_data['message'] = '내용이 300자를 초과했습니다.'
        return jsonify(res_data), 400
    print('checkpoint3')
    app.db.create_article(payload)
    print('checkpoint4')
    res_data['status'] = True
    return jsonify(res_data), 200


# 팔로우 기능
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


# 언팔로우
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


# 유저 정보 얻기
@api_miniter.route('/user/<int:user_id>', methods=['GET'])
def get_user(user_id: int):
    user = SIMPLE_DB['users'].get(user_id)
    if not user:
        return "No such user", 400

    return jsonify(user), 200
