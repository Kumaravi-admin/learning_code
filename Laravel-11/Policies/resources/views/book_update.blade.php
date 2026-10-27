<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<h1>Update Book Reocde</h1>

<form action="{{route('books.update',$book->id)}}" method="POST">
    @csrf
    @method('PUT')
    <label for="title">Book Title</label>
    <input type="text" id="title" name="title" value="{{$book->title}}">
    <label for="price">Book Price</label>
    <input type="text" id="price" name="price" value="{{$book->price}}">
    <button type="submit">Save</button>
</form>
</body>
</html>