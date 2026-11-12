// Thin JSON client for /api/v1 (Rails: app/javascript/lib/fetch.js).
//
// The JWT is stored in an httpOnly cookie, so requests rely on
// `withCredentials` instead of an Authorization header. The CSRF token is
// mirrored in a readable cookie and echoed back in a header (double submit).
import axios from 'axios';

const CSRF_COOKIE = 'csrf_token';
const CSRF_HEADER = 'X-CSRF-Token';
const SAFE_METHODS = ['GET', 'HEAD', 'OPTIONS'];

function readCookie(name) {
    return document.cookie.split('; ').reduce((found, part) => {
        const separator = part.indexOf('=');
        if (separator === -1) {
            return found;
        }
        const key = part.slice(0, separator);
        return key === name ? decodeURIComponent(part.slice(separator + 1)) : found;
    }, '');
}

const api = axios.create({
    baseURL: '/api/v1',
    withCredentials: true,
    headers: { Accept: 'application/json' },
});

api.interceptors.request.use((config) => {
    const method = (config.method || 'get').toUpperCase();
    if (!SAFE_METHODS.includes(method)) {
        const token = readCookie(CSRF_COOKIE);
        if (token) {
            config.headers[CSRF_HEADER] = token;
        }
    }
    return config;
});

// The API answers `{status, data, error_message}`; unwrap it once so views stay small.
export function payload(response) {
    return response?.data?.data ?? null;
}

export function errorMessage(error, fallback = 'Something went wrong') {
    return error?.response?.data?.error_message || error?.message || fallback;
}

export default api;
