from flask.json import JSONEncoder


class DataEncoder(JSONEncoder):
    """ Flask의 json 인코더의 변환 타입을 확장
    - JSON 통신에 필요한 타입이 생길 때마다 추가해준다.
    - jsonify()에 이용된다.
    """
    def default(self, o):
        # list는 json에서 지원하는 객체
        if isinstance(o, set):
            return list(set)

        return super().default(o)
