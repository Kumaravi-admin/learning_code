<?php

use App\Http\Controllers\EmailController;
use Illuminate\Support\Facades\Route;

Route::get('/contact', [EmailController::class, 'ContactForm']);
Route::post('/send-email', [EmailController::class, 'sendContactEmail'])->name('sendmail');
