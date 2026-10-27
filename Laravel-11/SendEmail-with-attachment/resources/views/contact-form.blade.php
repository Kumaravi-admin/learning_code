<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <div>
        <form action="{{route('sendmail')}}" method="POST" enctype="multipart/form-data">
        @csrf
        <label for="name">Name</label>
        <input type="text" name="name" id="name" placeholder="Your Name">

        <label for="email">Email Address</label>
        <input type="email" name="email" id="email" placeholder="Your Email">

        <label for="subject">Subject</label>
        <input type="text" name="subject" id="subject" placeholder="Subject">

        <label for="message">Message</label>
        <input type="text" name="message" id="message" placeholder="Message">

        <label for="file">File</label>
        <input type="file" name="attachment" id="file">

        <input type="submit" value="Submit">

        </form>
    </div>
    
</body>
</html>