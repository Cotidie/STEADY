from app import create_app


app = create_app(debug=True)
app.run(host='0.0.0.0')
