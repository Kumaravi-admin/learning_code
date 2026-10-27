<?php

namespace App\Http\Controllers;

use App\Mail\Welcomeemail;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Mail;

class EmailController extends Controller
{

    public function ContactForm()
    {
        return view('contact-form');
    }

    public function sendContactEmail(Request $request)
    {
        $request->validate([
            'name' => 'required',
            'email' => 'required|email',
            'subject' => 'required|min:5|max:100',
            'message' => 'required|min:5|max:255',
            'attachment' => 'required|mimes:pdf,doc,docx,xls,xlsx|max:2048',

        ]);
        $fileName = time() . "." . $request->file('attachment')->extension();

        $request->file('attachment')->move('uploads', $fileName);


        $adminEmail = 'anjeshravi8@yahoo.com';
        $response = Mail::to($adminEmail)->send(new Welcomeemail($request->all(), $fileName));

        if ($response) {
            return back()->with('success', 'Thank you for contact us');
        } else {
            return back()->with('error', 'unable to submit form data. Please try again');
        }



    }

    // // public function SendEmail()
    // // {
    // //     $toEmail = 'anjeshravi8@yahoo.com';
    // //     $message = 'Sending mail from Laravel Project';
    // //     $subject = 'Test Email Send Throw Laravel';
    // //     $details = [
    // //         'name' => 'Yahoo Baba',
    // //         'product' => 'Laravel',
    // //         'price' => 0
    // //     ];

    // //     $request = Mail::to($toEmail)->send(new Welcomeemail($message, $subject, $details));
    // //     // Mail::to($toEmail)->cc($mailid)->send(new Welcomeemail($message, $subject, $details));

    // //     // Mail::to($toEmail)->bcc($blindCorbonCopyMailId)->send(new Welcomeemail($message, $subject, $details));

    // //     dd($request);
    // }
}
