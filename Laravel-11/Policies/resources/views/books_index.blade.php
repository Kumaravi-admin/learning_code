<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h1>Books Page</h1>

    <h3>
        @if (Auth::check())
            {{Auth::user()->name}}
        @endif
    </h3>
        <a href="{{route('books.create')}}">Create Book</a>
        <a href="{{route('logout')}}">Logout</a>
    <table>
        <tr>
            <th>id</th>
            <th>Title</th>
            <th>Price</th>
            <th>User ID</th>
    
        </tr>
        @foreach ($Books as $Book)
        <tr>
            <td>{{$Book->id}}</td>
            <td>{{$Book->title}}</td>
            <td>{{$Book->price}}</td>
            <td>{{$Book->user_id}}</td>
            <td>
            {{-- @if (Auth::user()->can('update',$Book))
                   You are Authrized.
            @endif  --}}

               

            @can('update', $Book)
             {{-- @canany(['update','view'], $Book) --}}
            <a href="{{route('books.show',$Book->id)}}">View</a>
            <a href="{{route('books.edit',$Book->id)}}">Update</a>
            <a href="">Delete</a></td>
            @else
            <h6>You are not Authrized.</h6>
            @endcan

            @unless (Auth::user()->can('update',$Book))
                You are not Authrized
            @endunless
            
        </tr>
    
        @endforeach
    
    </table>
</body>
</html>