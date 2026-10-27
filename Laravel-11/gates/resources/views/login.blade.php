<h1>Login</h1>
<form action="{{route('loginMatch')}}" method="POST">
    @csrf

    <label for="useremal">Email</label>
    <input type="text" id="useremail" name="email">

    <label for="userpassword">Password</label>
    <input type="password" id="userpassword" name="password">

    <button type="submit">Login</button>

</form>

@if ($errors->any())
<ul>
    @foreach ($errors->all() as $error)
    <li>{{$error}}</li>
        
    @endforeach
</ul>
    
@endif