<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h1>Login Page</h1>
    <div>
        <label for="email">Enter your email</label>
        <input type="email" id="email" name="email">
        <label for="password">Enter your password</label>
        <input type="password" id="password" name="password">
        <button id="button">Login</button>

    </div>

    <script src="https://code.jquery.com/jquery-3.7.1.min.js" integrity="sha256-/JqT3SQfawRcv/BIHPThkBvs0OEvtFFmqPF/lYI/Cxo=" crossorigin="anonymous"></script>

    <script>
        $(document).ready(function(){
            $('#button').on('click',function(){

                const email = $('#email').val(); 
                const password = $('#password').val(); 

                $.ajax({
                url: 'api/login',
                type: 'POST',
                contentType:'application/json',
                data: JSON.stringify({
                    email:email,
                    password:password
                }),
               success:function(responce){
                console.log(responce);
                
                localStorage.setItem('api_token',responce.token);
                window.location.href = 'http://localhost:8000/allpost';

               },
               error:function(xhr,status,error){
                alert('Error: ' + xhr.responseText);

               }
            });

        });
    });
    </script>
</body>
</html>