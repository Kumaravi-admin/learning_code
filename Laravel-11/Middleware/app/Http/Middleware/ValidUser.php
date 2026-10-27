<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Symfony\Component\HttpFoundation\Response;

class ValidUser
{

    public function handle(Request $request, Closure $next): Response
    {
        echo '<h1>now you are in Valid User Page</h1>';
        if (Auth::check()) {
            return $next($request);
        } else {
            return redirect()->route('loginPage');
        }
    }

    // public function handle(Request $request, Closure $next, string $role): Response
    // {
    //     echo '<h1>now you are in Valid User Page</h1>';
    //     if (Auth::check() && Auth::user()->role == $role) {
    //         return $next($request);
    //     } else {
    //         return redirect()->route('loginPage');
    //     }
    // }
    public function terminate(Request $request, Response $response): void
    {
        echo '<h1>Terminating the User</h1>';
    }
}
