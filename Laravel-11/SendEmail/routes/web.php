<?php

use App\Http\Controllers\EmailController;
use Illuminate\Support\Facades\Route;

Route::get('/send-email', [EmailController::class, 'SendEmail']);
