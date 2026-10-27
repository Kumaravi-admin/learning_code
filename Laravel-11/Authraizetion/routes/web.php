<?php

use App\Http\Controllers\UserController;
use Illuminate\Support\Facades\Route;

Route::get('/home', [UserController::class, 'homePage']);

Route::view('register', 'register')->name('register');
Route::post('/registerSave', [UserController::class, 'register'])->name('register.save');

Route::view('login', 'login')->name('login');
Route::post('/loginMatch', [UserController::class, 'login'])->name('loginMatch');

Route::get('/dashboard', [UserController::class, 'dashboardPage'])->name('dashboard');
Route::get('/inner', [UserController::class, 'innerPage'])->name('innerPage');

Route::get('/logout', [UserController::class, 'logout'])->name('logout');