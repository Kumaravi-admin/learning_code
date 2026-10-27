<?php

use App\Http\Controllers\UserController;
use Illuminate\Support\Facades\Route;

Route::get('/user/add', [UserController::class, 'addUser']);
Route::get('/user/show', [UserController::class, 'showUser'])->name('user.show');