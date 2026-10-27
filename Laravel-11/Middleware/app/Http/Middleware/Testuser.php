<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class Testuser
{
    public function handle(Request $request, Closure $next): Response
    {

        echo '<h1>you are in Valid TestUser Middleware</h1>';

        return $next($request);
    }
}
