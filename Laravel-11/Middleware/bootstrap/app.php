<?php

use App\Http\Middleware\Testuser;
use App\Http\Middleware\ValidUser;
use Illuminate\Foundation\Application;
use Illuminate\Foundation\Configuration\Exceptions;
use Illuminate\Foundation\Configuration\Middleware;

return Application::configure(basePath: dirname(__DIR__))
    ->withRouting(
        web: __DIR__ . '/../routes/web.php',
        commands: __DIR__ . '/../routes/console.php',
        health: '/up',
    )
    ->withMiddleware(function (Middleware $middleware) {
        $middleware->alias([
            'testMiddleware' => Testuser::class,
            'ValidUser' => ValidUser::class
        ]);
        $middleware->appendToGroup('ok-user', [
            ValidUser::class,
            Testuser::class
        ]);

        // $middleware->append(TestUser::class);
        // $middleware->use([
        //     TestUser::class,
        //     ValidUser::class
        // ]);
    })
    ->withExceptions(function (Exceptions $exceptions) {
        //
    })->create();
