from app import create_app


app = create_app(debug=False)
print(app.config)
app.run(host='0.0.0.0')
