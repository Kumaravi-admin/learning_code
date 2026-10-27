<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h1>Post Page</h1>

    <table>
        <tr>
            <th>id</th>
            <th>title</th>
            <th>description</th>
            <th>user_id</th>
            <th>Update</th>
            <th>Delete</th>
        </tr>
            @foreach ($posts as $post)
            <tr>
                <td>{{$post->id}}</td>
                <td>{{$post->title}}</td>
                <td>{{$post->description}}</td>
                <td>{{$post->user_id}}</td>
                <td><a href="{{route('post.update',$post->id)}}">Update</a></td>
            </tr>
                
            @endforeach
    </table>
    
</body>
</html>