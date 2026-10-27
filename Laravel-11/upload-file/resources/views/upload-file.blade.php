'<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <div>
        <h2>Upload File</h2>
        <form action="{{route('user.store')}}" method="POST" enctype="multipart/form-data">
            @csrf
            <input type="file" name="photo" accept=".png,.jpeg,.jpg"> 
        <input type="submit">
        </form>
        
    </div>
    <div>
        @foreach ($users as $user)
        <div>
        <img src="{{asset('/storage/images/'.$user->file_name)}}" style="height: auto;width:200px">
         <form action="{{route('user.destroy',$user->id)}}" method="POST">
           @method('DELETE')
           @csrf
           <button type="submit" style="background-color: red">DELETE</button>
           <a href="{{route('user.edit',$user->id)}}">Update</a>     
        </form>   
    </div>   
        @endforeach
    </div>

</body>
</html>'