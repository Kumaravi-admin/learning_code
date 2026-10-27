<?php

use App\Http\Controllers\BookController;
use App\Http\Controllers\UserController;
use GuzzleHttp\Middleware;
use Illuminate\Support\Facades\Route;

Route::get('/', [UserController::class, 'loginPage'])->name('home');
Route::post('/login', [UserController::class, 'login'])->name('login');
Route::post('/logout', [UserController::class, 'logout'])->name('logout');

Route::resource('books', BookController::class)->middleware('auth');

// Route::put('/book/update/{id}', [BookController::class, 'update'])
//     // ->can('update','view');
//     ->Middleware('can:update', 'view');