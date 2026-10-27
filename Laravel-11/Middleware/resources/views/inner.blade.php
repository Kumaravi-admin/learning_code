<h1>Inner Page</h1>

@if (auth()->check())
    {{Auth::user()->name}}
@endif