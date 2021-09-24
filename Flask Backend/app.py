from flask import Flask
from flask_cors import CORS
from blueprint import api_main, api_miniter


# 팩토리 패턴
def create_app(debug=False):
    app = Flask(__name__)

    # Blueprints 등록
    app.register_blueprint(api_main)
    app.register_blueprint(api_miniter, url_prefix='/minister')

    # config 등록
    app.config.from_pyfile('config.py')
    app.config.update(DEBUG=debug)

    # CORS 세팅
    CORS(app, resources={r'/*': {'origins': '*'}})

    # 메인페이지
    @app.route("/")
    def hello_world():
        return "HEllo!"

    return app