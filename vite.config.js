import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import tailwindcss from '@tailwindcss/vite';
import path from 'node:path';

// Rails-like asset pipeline driven by the mode instead of NODE_ENV:
//   development / test -> static/dev  (served from disk by Fiber at /public/dev)
//   production         -> static/dist (embedded into the Go binary, /public/dist)
// Keep in sync with app/views/assets.go (Configure).
const ROOT_DIR = process.cwd();

const OUTPUT_DIRECTORIES = {
    development: 'static/dev',
    test: 'static/dev',
    production: 'static/dist',
};

export default defineConfig(({ mode }) => {
    const isProduction = mode === 'production';
    const outDir = OUTPUT_DIRECTORIES[mode] ?? OUTPUT_DIRECTORIES.development;

    return {
        // Assets are served by Go, so Vite must not copy a public directory.
        publicDir: false,
        plugins: [
            vue(),
            // Tailwind CSS v4 runs inside Vite (replaces postcss + tailwindcss CLI).
            tailwindcss(),
        ],
        resolve: {
            alias: {
                // app/javascript is the application; static/javascript keeps the
                // reusable components that Storybook still imports.
                '@': path.resolve(ROOT_DIR, 'app/javascript'),
                '@legacy': path.resolve(ROOT_DIR, 'static/javascript'),
            },
        },
        build: {
            outDir,
            emptyOutDir: true,
            // The Go shell reads .vite/manifest.json to link the hashed files.
            manifest: true,
            sourcemap: !isProduction,
            minify: isProduction ? 'esbuild' : false,
            target: isProduction ? 'es2019' : 'esnext',
            reportCompressedSize: false,
            rollupOptions: {
                input: {
                    main: path.resolve(ROOT_DIR, 'app/javascript/main.js'),
                },
                output: {
                    // Hashed filenames -> safe `immutable` caching (app/views/handler.go).
                    entryFileNames: 'javascript/[name]-[hash].js',
                    chunkFileNames: 'javascript/[name]-[hash].js',
                    assetFileNames: 'assets/[name]-[hash][extname]',
                },
            },
        },
    };
});
