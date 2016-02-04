python -m json.tool gotest.guess > got.guess
diff got.guess want.guess
