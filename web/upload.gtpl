<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>upload</title>
</head>
<body>
<form action="/upload" enctype="multipart/form-data" method="post">
    <input type="file" name="upFile">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="上传">
</form>
</body>
</html>