<h1>Register</h1>

<form action="{{route('register.save')}}" method="POST">
    @csrf
    <label for="username">Name</label>
    <input type="text" id="username" name="name">

    <label for="useremal">Email</label>
    <input type="text" id="useremail" name="email">

    <label for="userpassword">Password</label>
    <input type="password" id="userpassword" name="password">

    <label for="password-confirm">Confirm Password</label>
    <input type="password" id="password-confirm" name="password_confirmation">
    <button type="submit">Regester</button>

</form>

@if ($errors->any())
<ul>
    @foreach ($errors->all() as $error)
    <li>{{$error}}</li>
        
    @endforeach
</ul>
    
@endif