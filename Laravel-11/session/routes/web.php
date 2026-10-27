
<?php

use App\Http\Controllers\SessionController;
use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('welcome');
});
Route::get('/session/show', [SessionController::class, 'show_session']);
Route::get('/session/store', [SessionController::class, 'store_session']);
Route::get('/session/delete', [SessionController::class, 'delete_session']);


