from flask  import Blueprint, request, jsonify

api_main = Blueprint('test', import_name=__name__)


@api_main.route('/add', methods=['GET', 'POST'])
def add_highlight():
    if request.method == 'GET':
        req_data = request.args
    else:
        req_data = request.get_json()
    print(req_data)
    # 데이터 받기
    user = req_data.get('user')
    highlight = req_data.get('highlight')

    # DB에 저장한다.
    # DB 저장 코드

    if request.method == 'GET':
        return f"저장에 성공: {highlight}"
    else:
        return jsonify({
            'status': '성공',
            'value': highlight
        })

