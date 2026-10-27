{{-- @if (Gate::allows('isAdmin'))
<a href="" style="color: green">Admin Pannel</a>
@endif --}}


@can('isAdmin')
<a href="" style="color: green">Admin Pannel</a>
@else
<h1>You Are Not Admin</h1>
@endcan

@cannot('isAdamin')
<a href="" style="color: green">Other Pannel</a>    
@endcannot


<a href="{{route('user.profile',Auth::id())}}" style="color: rgb(66, 27, 139)">Profile</a>
<a href="{{route('post.show',Auth::id())}}" style="color: chocolate">Post</a>
<a href="{{route('logout')}}" style="color: red">Logout</a>