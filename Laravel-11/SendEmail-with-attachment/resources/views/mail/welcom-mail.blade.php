<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>{{$request['subject']}}</title>
</head>
<body>
    <h3>Hello Admin.</h3>
    <p>Name : {{$request['name']}}</p>
    <p>Email : {{$request['email']}}</p>
    <p>Subject : {{$request['subject']}}</p>
    <p>Message : {{$request['message']}}</p>

    {{-- if you want show image in mail message use super global variable $message & embed() method in embed() method give image apslute path --}}
    <img src="{{$message->embed('image_path')}}" alt="">
    
</body>
</html>