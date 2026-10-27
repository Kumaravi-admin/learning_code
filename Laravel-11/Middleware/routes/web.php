<?php

use App\Http\Controllers\UserController;
use App\Http\Middleware\Testuser;
use App\Http\Middleware\ValidUser;
use Illuminate\Support\Facades\Route;




Route::controller(UserController::class)->group(function () {
    Route::get('/', 'homePage')->name('homePage');
    Route::get('/login', 'loginPage')->name('loginPage');
    Route::get('/register', 'registerPage')->name('registerPage');
    Route::post('/registersave', 'register')->name('register.save');
    Route::post('/loginmatch', 'login')->name('loginMatch');
    Route::get('/logout', 'logout')->name('logout');
});

// Route::middleware([ValidUser::class, Testuser::class])->group(function () {
//     Route::get('/dashboard', 'dashboardPage')->name('dashboard');
//     Route::get('/dashboard/inner', 'innerPage')->name('innerPage')->withoutMiddleware(Testuser::class);
// });

// Route::withoutMiddleware(Testuser::class)->group(function () {
//     Route::get('/dashboard/inner', 'innerPage')->name('innerPage');
// });


// Route::middleware(['ok-user'])->group(function () {
//     Route::get('/dashboard', [UserController::class, 'dashboardPage'])->name('dashboard');
//     Route::get('/dashboard/inner', [UserController::class, 'innerPage'])->name('innerPage');
// });


// Route::middleware(['ValidUser:admin', Testuser::class])->group(function () {
//     Route::get('/dashboard', [UserController::class, 'dashboardPage'])->name('dashboard');
//     Route::get('/dashboard/inner', [UserController::class, 'innerPage'])->name('innerPage');
// });

Route::middleware('ValidUser')->group(function () {
    Route::get('/dashboard', [UserController::class, 'dashboardPage'])->name('dashboard');
    Route::get('/dashboard/inner', [UserController::class, 'innerPage'])->name('innerPage');
});