export function saveToLocalStorage(key, value) {
  try {
    localStorage.setItem(key, JSON.stringify(value));
  } catch (e) {
    console.error("Error saving to localStorage", e);
  }
}

export function getFromLocalStorage(key) {
  try {
    const value = localStorage.getItem(key);
    return value ? JSON.parse(value) : null;
  } catch (e) {
    console.error("Error getting data from localStorage", e);
    return null;
  }
}

export function removeFromLocalStorage(key) {
  try {
    localStorage.removeItem(key);
  } catch (e) {
    console.error("Error removing data from localStorage", e);
  }
}

export function clearLocalStorage() {
  try {
    localStorage.clear();
  } catch (e) {
    console.error("Error clearing localStorage", e);
  }
}
