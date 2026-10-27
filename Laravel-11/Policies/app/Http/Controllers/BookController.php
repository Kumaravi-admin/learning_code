<?php

namespace App\Http\Controllers;

use App\Models\Book;
use Auth;
use Gate;
use Illuminate\Http\Request;

class BookController extends Controller
{
    public function index()
    {
        $Books = Book::all();
        return view('books_index', compact('Books'));

    }

    public function create()
    {
        return view('books_create');
    }
    public function store(Request $request)
    {
        $request->validate([
            'title' => 'required',
            'price' => 'required|numeric'
        ]);

        $user_id = Auth::id();

        Book::create([
            'title' => $request->title,
            'price' => $request->price,
            'user_id' => $user_id
        ]);
        return redirect()->route('books.index');
    }

    public function show(Book $book)
    {
        Gate::authorize('view', $book);
        return view('book_show', compact('book'));

    }

    public function edit(Book $book)
    {
        Gate::authorize('edit', $book);
        $book = Book::findOrFail($book->id);
        return view('book_update', compact('book'));
    }

    public function update(Request $request, Book $book)
    {
        // Gate::authorize('update', $book);
        if ($request->user()->cannot('update', $book)) {
            abort(403, 'You are not Authrized');
        }


        $request->validate([
            'title' => 'required',
            'price' => 'required|numeric'
        ]);

        Book::where('id', $book->id)->update([
            'title' => $request->title,
            'price' => $request->price

        ]);
    }

}
