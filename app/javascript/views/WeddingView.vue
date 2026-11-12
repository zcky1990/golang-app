<template>
    <main class="min-h-screen bg-slate-50 px-4 py-10">
        <div class="mx-auto max-w-3xl">
            <header class="mb-8">
                <p class="font-dosis text-xs uppercase tracking-[0.3em] text-slate-500">Wedding</p>
                <h1 class="font-playfair-display text-3xl text-slate-900">Invitation</h1>
            </header>

            <p v-if="loading" class="font-dosis text-slate-600">Loading…</p>

            <div v-else-if="error" class="rounded-lg border border-red-200 bg-red-50 p-4">
                <p class="font-dosis text-sm text-red-700">{{ error }}</p>
                <button
                    type="button"
                    class="mt-3 rounded-md bg-red-600 px-3 py-1.5 text-sm font-medium text-white"
                    @click="load"
                >
                    Try again
                </button>
            </div>

            <pre v-else class="overflow-x-auto rounded-lg bg-white p-4 text-xs text-slate-700 shadow-sm">{{ wedding }}</pre>
        </div>
    </main>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

import api, { errorMessage, payload } from '../lib/api';

const route = useRoute();

const wedding = ref(null);
const loading = ref(true);
const error = ref('');

async function load() {
    loading.value = true;
    error.value = '';

    try {
        // The id comes from the URL (/wedding/:id), validated by the Go router,
        // and again here so a bad link fails fast.
        const response = await api.get(`/wedding/${encodeURIComponent(route.params.id)}`);
        wedding.value = payload(response);
    } catch (requestError) {
        error.value = errorMessage(requestError, 'Unable to load the invitation');
    } finally {
        loading.value = false;
    }
}

onMounted(load);
</script>
