<html>
<head>
    <title></title>
</head>
<body>
<form action="/login" method="post">
    用户名：<input type="text" name="username">
    密码： <input type="password" name="password">
    水果： <select name="fruit">
            <option value="apple">苹果</option>
            <option value="pear">桃子</option>
            <option value="banana">香蕉</option>
          </select>
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="登陆">
</form>
</body>
</html>