from flask import Flask

app = Flask(__name__)

class Fibonacci:
  def __init__(self) -> None:
    self.current_fib: int = 0
    self.next_fib: int = 1
  
  def current(self) -> int:
    return self.current_fib
  
  def next(self) -> int:
    res = self.next_fib
    new_fib = self.next_fib + self.current_fib
    self.current_fib = self.next_fib
    self.next_fib = new_fib
    return res

  def previous(self) -> int or str:
    if self.current_fib == 0:
      raise Exception("Cannot go previous")

    temp = self.next_fib - self.current_fib
    self.next_fib = self.current_fib
    self.current_fib = temp

    return self.current_fib

fib_generator = Fibonacci()

@app.route('/')
def index():
  response = "<a href=\"/previous\">/previous</a> - For previous fibonacci number\n</br>" + \
		"<a href=\"/current\">/current</a> - For current fibonacci number\n</br>" + \
		"<a href=\"/next\">/next</a> - For next fibonacci number\n</br>"
  return response

@app.route('/next')
def next():
  return str(fib_generator.next())

@app.route('/previous')
def previous():
  try:
    return str(fib_generator.previous())
  except Exception as e:
    return str(e)

@app.route('/current')
def current():
  return str(fib_generator.current())


if __name__ == '__main__':
  app.run(host='0.0.0.0', port=3035)