# Third Party
from flask import Flask
from flask_cors import CORS

# Custom
from blueprint import api_main, api_miniter
from database  import DataBase


# 팩토리 패턴
def create_app(debug=False):
    app = Flask(__name__)

    # Blueprints 등록
    app.register_blueprint(api_main)
    app.register_blueprint(api_miniter, url_prefix='/minister')

    # config 등록
    app.config.from_pyfile('config.py')
    app.config.update(DEBUG=debug)

    # 데이터베이스 등록
    app.config.db = app.config.get_namespace('MONGODB_')
    app.db = DataBase(app.config.db.get('host'), app.config.db.get('db'))

    # CORS 세팅
    CORS(app, resources={r'/*': {'origins': '*'}})

    # 메인페이지
    @app.route("/")
    def hello_world():
        return "HEllo!"

    return app