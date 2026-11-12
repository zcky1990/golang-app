<template>
    <main class="flex min-h-screen items-center justify-center bg-slate-50 px-6">
        <div class="text-center">
            <p class="font-dosis text-sm uppercase tracking-[0.3em] text-slate-500">404</p>
            <h1 class="mt-2 font-playfair-display text-3xl text-slate-900">Page not found</h1>
            <p class="mt-3 font-dosis text-slate-600">
                The page you are looking for does not exist.
            </p>
            <RouterLink to="/" class="mt-6 inline-block rounded-md bg-slate-900 px-4 py-2 text-sm font-medium text-white">
                Back to home
            </RouterLink>
        </div>
    </main>
</template>

<script setup>
import { RouterLink } from 'vue-router';
</script>
