<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<h1>Creat New Book Recode</h1>

<form action="{{route('books.store')}}" method="POST">
    @csrf
    
    <label for="title">Book Title</label>
    <input type="text" id="title" name="title">
    <label for="price">Book Price</label>
    <input type="text" id="price" name="price">
    <button type="submit">Save</button>
</form>
</body>
</html>