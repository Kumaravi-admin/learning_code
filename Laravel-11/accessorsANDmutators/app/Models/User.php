<?php

namespace App\Models;

use Illuminate\Support\Number;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Casts\Attribute;
use Illuminate\Database\Eloquent\Factories\HasFactory;

class User extends Model
{
    use HasFactory;
    public $timestamps = false;
    protected $guarded = [];

    public function setEmailAttribute($value)
    {
        $this->attributes['email'] = strtolower($value);
    }
    // public function setUserNameAttribute($value)
    // {
    //     $this->attributes['user_name'] = strtolower($value);
    // }
    public function setPasswordAttribute($value)
    {
        $this->attributes['password'] = bcrypt($value);
    }
    public function getDobAttribute($value)
    {
        return date('D M Y', strtotime($value));
    }
    // public function getUserNameAttribute($value)
    // {
    //     return ucwords($value);
    // }
    public function getSalaryAttribute($value)
    {
        return Number::currency($value, in: 'INR');
    }
    protected function UserName(): Attribute
    {
        return Attribute::make(
            get: fn(string $value) => ucwords($value),
            set: fn(string $value) => strtolower($value)

        );
    }

}
