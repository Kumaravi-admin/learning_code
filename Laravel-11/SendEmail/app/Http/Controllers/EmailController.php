<?php

namespace App\Http\Controllers;

use App\Mail\Welcomeemail;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Mail;

class EmailController extends Controller
{
    public function SendEmail()
    {
        $toEmail = 'anjeshravi8@yahoo.com';
        $message = 'Sending mail from Laravel Project';
        $subject = 'Test Email Send Throw Laravel';
        $details = [
            'name' => 'Yahoo Baba',
            'product' => 'Laravel',
            'price' => 0
        ];

        $request = Mail::to($toEmail)->send(new Welcomeemail($message, $subject, $details));
        // Mail::to($toEmail)->cc($mailid)->send(new Welcomeemail($message, $subject, $details));

        // Mail::to($toEmail)->bcc($blindCorbonCopyMailId)->send(new Welcomeemail($message, $subject, $details));

        dd($request);
    }
}
