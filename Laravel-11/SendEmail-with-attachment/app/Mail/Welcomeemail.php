<?php

namespace App\Mail;

use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Mail\Mailable;
use Illuminate\Mail\Mailables\Attachment;
use Illuminate\Mail\Mailables\Content;
use Illuminate\Mail\Mailables\Envelope;
use Illuminate\Queue\SerializesModels;

class Welcomeemail extends Mailable
{
    use Queueable, SerializesModels;

    public $request;
    public $fileName;

    public function __construct($request, $fileName)
    {
        $this->request = $request;
        $this->fileName = $fileName;


    }

    /**
     * Get the message envelope.
     */
    public function envelope(): Envelope
    {
        return new Envelope(
            subject: 'Contact Form',
        );
    }

    /**
     * Get the message content definition.
     */
    public function content(): Content
    {
        return new Content(
            view: 'mail.welcom-mail',
        );
    }

    /**
     * Get the attachments for the message.
     *
     * @return array<int, \Illuminate\Mail\Mailables\Attachment>
     */
    public function attachments(): array
    {
        $attachment = [];

        if ($this->fileName) {
            $attachment = [
                Attachment::fromPath(public_path('/uploads/' . $this->fileName))
            ];
        }
        return $attachment;
    }
}
// Attachment::fromStorage('file_path')

// Attachment::fromStorageDisk('s3','file_path');