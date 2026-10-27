<?php

use App\Http\Controllers\UserController;
use App\Http\Middleware\ValidateUser;
use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome');
});

Route::get('/login', [UserController::class, 'loginPage'])->name('loginPage');
Route::post('/loginMatch', [UserController::class, 'loginMatch'])->name('loginMatch');

// Route::get('/dashboard', [UserController::class, 'dashboardPage'])->name('dashboard')->middleware('can:isAdmin');

Route::get('/dashboard', [UserController::class, 'dashboardPage'])
    ->name('dashboard')->can('isAdmin');

Route::get('/profile/{id}', [UserController::class, 'profilePage'])->name('user.profile')->can('isAdmin');

Route::get('/post/{id}', [UserController::class, 'showPost'])->name('post.show');
Route::get('/single-post/{id}', [UserController::class, 'postUpdate'])->name('post.update');

Route::get('/logout', [UserController::class, 'logout'])->name('logout');
