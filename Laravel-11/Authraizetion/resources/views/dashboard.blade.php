<h1>Dashboard Page : Welcome {{Auth::user()->name}}</h1>
{{Auth::user()}}

<a href="{{route('innerPage')}}">Inner Page</a>
<a href="{{route('logout')}}">Logout</a>