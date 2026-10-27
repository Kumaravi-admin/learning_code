<?php

namespace App\Mail;

use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Mail\Mailable;
use Illuminate\Mail\Mailables\Content;
use Illuminate\Mail\Mailables\Envelope;
use Illuminate\Queue\SerializesModels;

class Welcomeemail extends Mailable
{
    use Queueable, SerializesModels;

    public $mailmessage;
    public $subject;
    public $details;

    protected $data;
    /**
     * Create a new message instance.
     */
    public function __construct($message, $subject, $details)
    {
        $this->mailmessage = $message;
        $this->subject = $subject;
        $this->details = $details;

    }

    /**
     * Get the message envelope.
     */
    public function envelope(): Envelope
    {
        return new Envelope(
            subject: $this->subject,
        );
    }

    /**
     * Get the message content definition.
     */
    public function content(): Content
    {
        return new Content(
            view: 'mail.welcom-mail',
            // text: 'hello how are you?',
            with: [
                'name' => $this->data['name'],
                'product' => $this->data['product'],
                'price' => $this->data['price'],
            ]
        );
    }

    /**
     * Get the attachments for the message.
     *
     * @return array<int, \Illuminate\Mail\Mailables\Attachment>
     */
    public function attachments(): array
    {
        return [];
    }
}
