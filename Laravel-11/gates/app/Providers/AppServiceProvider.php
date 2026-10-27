<?php

namespace App\Providers;

use App\Models\User;
use Illuminate\Support\Facades\Gate;
use Illuminate\Support\ServiceProvider;

class AppServiceProvider extends ServiceProvider
{

    public function boot(): void
    {
        Gate::define('isAdmin', function (User $user) {
            return $user->role === 'admin';
        });

        Gate::define('view-profile', function (User $user, $UserId) {
            return $user->id === $UserId;
        });

        Gate::define('update-post', function (User $user, $targetUser) {
            return $user->id === $targetUser;
        });

        Gate::before(function (User $user) {

        });
    }
}
