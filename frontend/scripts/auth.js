/* =====================
DOM-элементы
===================== */
const openModalBtn = document.getElementById('openAuthModal');
const closeModalBtn = document.getElementById('closeAuthModal');
const modalOverlay = document.getElementById('authModalOverlay');

const tabBtns = document.querySelectorAll('.tab-btn');
const tabContents = document.querySelectorAll('.tab-content');
const switchTabLinks = document.querySelectorAll('.switch-tab');

const registerForm = document.getElementById('registerForm');
const loginForm = document.getElementById('loginForm');

const logoutBtn = document.getElementById('logoutBtn');

/* =====================
Модальное окно
===================== */
function openModal() {
    modalOverlay.classList.add('active');
    document.body.style.overflow = 'hidden';
}

function closeModal() {
    modalOverlay.classList.remove('active');
    document.body.style.overflow = 'auto';
}

openModalBtn.addEventListener('click', openModal);
closeModalBtn.addEventListener('click', closeModal);

modalOverlay.addEventListener('click', (e) => {
    if (e.target === modalOverlay) closeModal();
});

document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape' && modalOverlay.classList.contains('active')) {
        closeModal();
    }
});

/* =====================
Переключение табов (Вход / Регистрация)
===================== */
function switchTab(tabId) {
    tabBtns.forEach(btn => btn.classList.toggle('active', btn.dataset.tab === tabId));
    tabContents.forEach(content => content.classList.toggle('active', content.id === tabId + 'Tab'));
}

tabBtns.forEach(btn => btn.addEventListener('click', () => switchTab(btn.dataset.tab)));

switchTabLinks.forEach(link => {
    link.addEventListener('click', (e) => {
        e.preventDefault();
        switchTab(link.dataset.tab);
    });
});

/* =====================
Регистрация
===================== */
registerForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const inputs = registerForm.querySelectorAll('input');
    const data = {
        first_name: inputs[0].value,
        last_name:  inputs[1].value,
        email:      inputs[2].value,
        password:   inputs[3].value,
        repeat_password: inputs[4].value
    };

    const response = await fetch('/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    });

    const result = await response.json();

    if (result.status === 'ok') {
        alert('Регистрация успешна');
        closeModal();
    } else {
        alert(result.error || 'Ошибка регистрации');
    }
});

/* =====================
Вход
===================== */
loginForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const inputs = loginForm.querySelectorAll('input');
    const data = {
        email:    inputs[0].value,
        password: inputs[1].value
    };

    const response = await fetch('/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    });

    const result = await response.json();

    if (result.status === 'ok') {
        await loadProfile();  // Обновляем профиль и кнопки входа/выхода
        closeModal();
    } else {
        alert(result.error || 'Ошибка входа');
    }
});

/* =====================
Профиль и логаут
===================== */
function showGuest() {
    document.querySelector('.name').textContent = 'Гость';
    document.querySelector('.email').textContent = '';
    document.querySelector('.profile').style.display = 'flex';
    openModalBtn.style.display = 'block';
    logoutBtn.style.display = 'none';
}

function showUser(user) {
    document.querySelector('.name').innerHTML = `${user.first_name}<br>${user.last_name}`;
    document.querySelector('.email').textContent = user.email;
    document.querySelector('.profile').style.display = 'flex';
    openModalBtn.style.display = 'none';
    logoutBtn.style.display = 'block';
}

async function loadProfile() {
    try {
        const response = await fetch('/profile');
        if (!response.ok) throw new Error();

        const user = await response.json();
        showUser(user);
    } catch {
        showGuest();
    }
}

// Обработка выхода
logoutBtn.addEventListener('click', async () => {
    try {
        await fetch('/logout', { method: 'POST' });
    } catch {
        // Игнорируем ошибку сети — всё равно очищаем интерфейс
    }
    showGuest();
});

// Проверка авторизации при загрузке страницы
window.addEventListener('load', loadProfile);