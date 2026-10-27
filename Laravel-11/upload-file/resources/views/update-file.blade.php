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
        <h2>Update File</h2>
        <form action="{{route('user.update',$user->id)}}" method="POST" enctype="multipart/form-data">
            @csrf
            @method('PUT')
            <img id="image" src="{{asset('/storage/images/'.$user->file_name)}}" style="height: auto;width:200px">
            <input type="file" onchange="document.querySelector('#image').src=window.URL.createObjectURL(this.files[0])" name="photo" accept=".png,.jpeg,.jpg"> 
        <button type="submit">Update</button>
        </form>
        
    </div>
    
</body>
</html>'