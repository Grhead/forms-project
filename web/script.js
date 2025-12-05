const API_BASE = "http://localhost:3000/api";
let section = "forms";
let selectedQuestions = [];
let newQuestionAnswers = [];
let allAvailableQuestions = [];
let currentListData = [];
let currentFormDetails = null;

function toggleMobileSidebar() {
    document.getElementById('sidebar').classList.toggle('mobile-open');
}

function showToast(title, message, type = 'info') {
    const container = document.getElementById('toastContainer');
    const toast = document.createElement('div');
    toast.className = `toast ${type}`;
    toast.innerHTML = `<div class="toast-title">${title}</div><div class="toast-msg">${message}</div>`;
    container.appendChild(toast);

    requestAnimationFrame(() => toast.classList.add('show'));

    setTimeout(() => {
        toast.classList.remove('show');
        setTimeout(() => toast.remove(), 300);
    }, 4000);
}

async function copyToClipboard(text) {
    try {
        await navigator.clipboard.writeText(text);
        showToast("Скопировано!", text, "success");
    } catch (err) {
        showToast("Ошибка", "Не удалось скопировать", "error");
    }
}

function formatReadableDate(dateString) {
    if (!dateString) return "Дата не указана";
    try {
        const date = new Date(dateString);
        return date.toLocaleDateString("ru-RU", {
            year: "numeric",
            month: "short",
            day: "numeric",
            hour: "2-digit",
            minute: "2-digit"
        });
    } catch (e) {
        return "Ошибка даты";
    }
}

function renderSkeleton(count = 3) {
    return Array(count).fill('<div class="skeleton h-40"></div>').join('');
}

async function apiGet(path) {
    try {
        const response = await fetch(API_BASE + path);
        if (!response.ok) throw new Error(response.status);
        return await response.json();
    } catch (error) {
        console.error("API Error", error);
        showToast("Ошибка сети", "Не удалось загрузить данные", "error");
        return null;
    }
}

async function apiPost(path, body) {
    try {
        const response = await fetch(API_BASE + path, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(body),
        });
        if (!response.ok) throw new Error(await response.text());
        const text = await response.text();
        return text ? JSON.parse(text) : {};
    } catch (error) {
        showToast("Ошибка сохранения", error.message, "error");
        throw error;
    }
}

function setSection(s) {
    section = s;
    document.querySelectorAll(".toolbar .btn").forEach(btn => btn.classList.remove("active"));
    const btn = document.getElementById(`btn${s.charAt(0).toUpperCase() + s.slice(1)}`);
    if (btn) btn.classList.add("active");

    document.getElementById('searchInput').value = '';
    document.getElementById('sidebar').classList.remove('mobile-open');
    currentFormDetails = null;
    updateUI();
}

async function updateUI() {
    const list = document.getElementById("list");
    const workspace = document.getElementById("workspace");

    list.innerHTML = renderSkeleton(5);
    workspace.innerHTML = '<div class="section-animate"><div class="skeleton h-40" style="width:50%"></div><div class="skeleton h-100"></div></div>';

    if (section === "forms") {
        await renderForms();
    } else if (section === "newQuestion") {
        await renderQuestions();
        renderNewQuestionForm();
    } else if (section === "newForm") {
        await renderQuestionListForForm();
        renderNewFormForm();
    }
}

function filterList(term) {
    term = term.toLowerCase();
    const filtered = currentListData.filter(item => {
        const title = (item.title || "").toLowerCase();
        const sub = (item.type || item.createdAt || "").toLowerCase();
        return title.includes(term) || sub.includes(term);
    });
    renderListItems(filtered);
}

function renderListItems(items) {
    const list = document.getElementById("list");
    if (items.length === 0) {
        list.innerHTML = '<div style="padding:20px; text-align:center; opacity:0.5">Ничего не найдено</div>';
        return;
    }
    list.innerHTML = items.map((item, idx) => {
        const delay = idx * 0.05;
        if (section === "forms") {
            return `
                    <div class="list-item" style="animation-delay: ${delay}s" onclick="openForm('${item.externalID}')">
                        <span class="form-title">${item.title || "Без названия"}</span>
                        <span class="form-meta">${formatReadableDate(item.createdAt)}</span>
                    </div>`;
        } else {
            let shortDesc = item.description || "";
            if (shortDesc.length > 15) {
                shortDesc = shortDesc.substring(0, 15) + "...";
            }

            const clickAction = section === "newForm" ? `handleQuestionClick(this)` : "";

            return `
                    <div class="list-item" style="animation-delay: ${delay}s" 
                         data-title="${item.title.replace(/"/g, '&quot;')}" 
                         onclick="${clickAction}">
                        <span class="form-title">${item.title}</span>
                        <span class="form-meta">
                            ${item.type} 
                            ${shortDesc ? `<span style="opacity:0.6; margin-left:6px; font-size:0.8em">(${shortDesc})</span>` : ''}
                        </span>
                    </div>`;
        }
    }).join("");
}

async function renderForms() {
    const forms = await apiGet("/forms");
    currentListData = Array.isArray(forms) ? forms : [];
    renderListItems(currentListData);
    document.getElementById("workspace").innerHTML = `
                <div class="section-animate" style="text-align:center; padding-top:50px; opacity:0.7">
                    <div style="font-size:3rem; margin-bottom:10px">📂</div>
                    <h2>Выберите форму</h2>
                    <p>Нажмите на форму в меню слева для просмотра деталей</p>
                </div>`;
}

async function renderQuestions() {
    const questions = await apiGet("/questions");
    currentListData = Array.isArray(questions) ? questions : [];
    renderListItems(currentListData);
}

async function renderQuestionListForForm() {
    await renderQuestions();
}

async function openForm(id) {
    const f = await apiGet("/form?form_id=" + id);
    if (!f || !f.externalID) return;

    currentFormDetails = f;

    const html = `
                <div class="section-animate">
                    <div style="display:flex; justify-content:space-between; align-items:flex-start">
                        <h2 class="section-title">${f.title}</h2>
                        <button class="btn btn-sm" onclick="copyToClipboard('${f.externalID}')">📋 ID</button>
                    </div>
                    
                    <div class="form-grid">
                        <div>ID:</div><div style="font-family:monospace; opacity:0.8">${f.externalID}</div>
                        <div>Документ:</div><div>${f.documentTitle}</div>
                        <div>Описание:</div><div>${f.description || "—"}</div>
                        <div>Дата:</div><div>${formatReadableDate(f.createdAt)}</div>
                    </div>

                    <button class="btn btn-success" 
                        onclick="confirmResubmit()">
                        🔄 Повторить форму
                    </button>

                    <h3 style="margin: 25px 0 15px;">Вопросы (${(f.questions || []).length})</h3>
                    ${(f.questions || []).map(q => `
                        <div class="question-box">
                            <div style="font-weight:600; font-size:1.1rem">${q.title}</div>
                            <div style="opacity:0.7; font-size:0.9rem; margin:4px 0">${q.description || ""}</div>
                            <div style="margin-top:8px">
                                <span style="background:rgba(255,255,255,0.1); padding:4px 8px; border-radius:6px; font-size:0.8rem">${q.type}</span>
                                ${q.isRequired ? '<span style="color:#ff9a9a; font-size:0.8rem; margin-left:5px">* Обязательный</span>' : ''}
                            </div>
                        </div>
                    `).join("")}
                </div>
            `;
    document.getElementById("workspace").innerHTML = html;
}

function renderNewQuestionForm() {
    document.getElementById("workspace").innerHTML = `
                <div class="section-animate">
                    <h2 class="section-title">Создать вопрос</h2>
                    <div class="form-grid">
                        <label>Заголовок</label>
                        <input class="glass-input" id="qTitle" placeholder="Например: Ваш возраст?">
                        
                        <label>Описание</label>
                        <textarea class="glass-textarea" id="qDesc" placeholder="Подсказка для пользователя..."></textarea>
                        
                        <label>Тип</label>
                        <select class="glass-input" id="qType" onchange="toggleAnswerOptions()">
                            <option value="TEXT">TEXT (Текст)</option>
                            <option value="RADIO">RADIO (Один выбор)</option>
                            <option value="CHECKBOX">CHECKBOX (Множество)</option>
                        </select>
                        
                        <label>Опции</label>
                        <label style="display:flex; align-items:center; gap:10px; cursor:pointer">
                            <input type="checkbox" id="qReq" style="width:20px; height:20px"> Обязательный вопрос
                        </label>
                    </div>

                    <div id="answerOptionsBox" style="display:none; background:rgba(255,255,255,0.05); padding:15px; border-radius:12px; margin-bottom:20px">
                        <h4 style="margin-bottom:10px">Варианты ответа</h4>
                        <div id="answerOptions"></div>
                        <button class="btn btn-sm" style="margin-top:10px" onclick="addAnswerOption()">+ Вариант</button>
                    </div>

                    <button class="btn btn-success" style="width:100%" onclick="submitNewQuestion()">💾 Сохранить вопрос</button>
                </div>
            `;
    newQuestionAnswers = [];
    toggleAnswerOptions();
}

function renderNewFormForm() {
    document.getElementById("workspace").innerHTML = `
                <div class="section-animate">
                    <h2 class="section-title">Конструктор формы</h2>
                    <div class="form-grid">
                        <label>Название</label> <input class="glass-input" id="newFormTitle">
                        <label>Документ</label> <input class="glass-input" id="newFormDoc" placeholder="Название файла выгрузки">
                        <label>Описание</label> <textarea class="glass-textarea" id="newFormDesc"></textarea>
                    </div>

                    <h3 style="margin-bottom:10px">Вопросы в форме</h3>
                    <p style="font-size:0.85rem; opacity:0.7; margin-bottom:15px">Нажмите на вопросы в списке слева, чтобы добавить их. Перетаскивайте блоки для сортировки.</p>
                    
                    <div id="selQ" class="question-list-dnd" style="min-height:100px; border:2px dashed rgba(255,255,255,0.1); border-radius:12px; padding:10px">
                        ${selectedQuestions.length === 0 ? '<div style="text-align:center; padding:30px; opacity:0.5">Список пуст</div>' : ''}
                    </div>

                    <button class="btn btn-success" style="width:100%; margin-top:20px" onclick="submitNewForm()">🚀 Опубликовать форму</button>
                </div>
            `;
    renderSelectedQuestions();
}

function toggleAnswerOptions() {
    const type = document.getElementById("qType").value;
    const box = document.getElementById("answerOptionsBox");
    const isSelect = (type === "RADIO" || type === "CHECKBOX");

    box.style.display = isSelect ? "block" : "none";
    if (isSelect && newQuestionAnswers.length === 0) newQuestionAnswers = ["", ""];
    if (isSelect) renderAnswerOptions();
}

function renderAnswerOptions() {
    document.getElementById("answerOptions").innerHTML = newQuestionAnswers.map((ans, i) => `
                <div style="display:flex; gap:10px; margin-bottom:8px">
                    <input class="glass-input" value="${ans.replace(/"/g, '&quot;')}" oninput="updateAnswerOption(${i}, this.value)" placeholder="Вариант ${i + 1}">
                    <button class="btn btn-sm" style="background:var(--red-accent); color:#fff" onclick="removeAnswerOption(${i})">✕</button>
                </div>
            `).join("");
}

function addAnswerOption() {
    newQuestionAnswers.push("");
    renderAnswerOptions();
}

function removeAnswerOption(i) {
    newQuestionAnswers.splice(i, 1);
    renderAnswerOptions();
}

function updateAnswerOption(i, val) {
    newQuestionAnswers[i] = val;
}

async function submitNewQuestion() {
    const title = document.getElementById("qTitle").value.trim();
    const type = document.getElementById("qType").value;
    const answers = newQuestionAnswers.filter(a => a.trim());
    const isSelectType = (type === "RADIO" || type === "CHECKBOX");

    if (!title) return showToast("Ошибка", "Введите заголовок", "error");

    if (isSelectType) {
        if (answers.length < 2) {
            return showToast("Ошибка", "Минимум 2 варианта ответа", "error");
        }
        const uniqueAnswers = new Set(answers.map(a => a.toLowerCase()));
        if (uniqueAnswers.size !== answers.length) {
            return showToast("Ошибка", "Варианты ответов должны быть уникальными (без учета регистра)", "error");
        }
    }


    try {
        await apiPost("/question", {
            title,
            description: document.getElementById("qDesc").value,
            type,
            isRequired: document.getElementById("qReq").checked,
            possibleAnswers: answers.map(content => ({
                content
            }))
        });

        showToast("Успех", "Вопрос создан", "success");
        setSection("newQuestion");
    } catch (e) { }
}

function handleQuestionClick(el) {
    const title = el.dataset.title.replace(/&quot;/g, '"');
    addQuestionToSelection(title);

    el.style.transform = "scale(0.98)";
    setTimeout(() => el.style.transform = "", 100);
}

function addQuestionToSelection(title) {
    const full = currentListData.find(q => q.title === title);

    if (full) {
        selectedQuestions.push(full);
        showToast("Добавлено", `Вопрос добавлен (${selectedQuestions.length})`, "success");
    }
    renderSelectedQuestions();
}

function renderSelectedQuestions() {
    const container = document.getElementById("selQ");
    if (!container) return;

    if (selectedQuestions.length === 0) {
        container.innerHTML = '<div style="text-align:center; padding:30px; opacity:0.5">Список пуст. Добавьте вопросы из меню.</div>';
        return;
    }

    container.innerHTML = selectedQuestions.map((q, index) => `
                <div class="question-box" draggable="true" data-index="${index}" data-title="${q.title.replace(/"/g, '&quot;')}" style="display:flex; align-items:center; justify-content:space-between">
                    <div style="display:flex; align-items:center">
                        <span class="dnd-handle">☰</span>
                        <div>
                            <div style="font-weight:600">${q.title}</div>
                            <small>${q.type}</small>
                        </div>
                    </div>
                    <button class="btn btn-sm" style="background:transparent; color:var(--red-accent); border:none" 
                        onclick="removeSelQ(${index})">Удалить</button>
                </div>
            `).join("");

    addDnDListeners();
}

function removeSelQ(index) {
    selectedQuestions.splice(index, 1);
    renderSelectedQuestions();
}

async function submitNewForm() {
    const title = document.getElementById("newFormTitle").value.trim();
    const doc = document.getElementById("newFormDoc").value.trim();

    if (!title || !doc) return showToast("Ошибка", "Заполните заголовок и название документа", "error");
    if (selectedQuestions.length === 0) return showToast("Внимание", "Добавьте хотя бы один вопрос", "error");

    try {
        await apiPost("/form", {
            title,
            documentTitle: doc,
            description: document.getElementById("newFormDesc").value,
            questions: selectedQuestions.map(q => q.title)
        });

        showToast("Готово!", "Форма успешно создана", "success");
        selectedQuestions = [];
        setSection("forms");
    } catch (e) { }
}

function confirmResubmit() {
    if (!currentFormDetails) return;

    const f = currentFormDetails;
    const qCount = (f.questions || []).length;

    const backdrop = document.getElementById('modalBackdrop');

    document.getElementById('modalTitle').textContent = "Повторить отправку?";
    document.getElementById('modalMessage').innerHTML = `Создать копию формы <b>"${f.title}"</b> с ${qCount} вопросами?`;

    const actions = document.getElementById('modalActions');
    actions.innerHTML = `
                <button class="btn" onclick="closeModal()">Отмена</button>
                <button class="btn btn-success" onclick="executeResubmit()">Подтвердить</button>
            `;

    backdrop.classList.add('open');
}

async function executeResubmit() {
    closeModal();
    if (!currentFormDetails) return;
    const f = currentFormDetails;

    try {
        const questions = (f.questions || []).map(q => q.title);
        await apiPost("/form", {
            title: f.title,
            documentTitle: f.documentTitle,
            description: f.description,
            questions
        });
        showToast("Успех", "Форма скопирована", "success");
        renderForms();
    } catch (e) { }
}

function closeModal() {
    document.getElementById('modalBackdrop').classList.remove('open');
}

let draggedItem = null;

function addDnDListeners() {
    const items = document.querySelectorAll('#selQ .question-box');
    const container = document.getElementById('selQ');

    items.forEach(item => {
        item.addEventListener('dragstart', () => {
            draggedItem = item;
            item.style.opacity = '0.5';
        });
        item.addEventListener('dragend', () => {
            draggedItem = null;
            item.style.opacity = '1';
        });
    });

    container.addEventListener('dragover', e => {
        e.preventDefault();
        const afterElement = getDragAfterElement(container, e.clientY);
        if (afterElement == null) {
            container.appendChild(draggedItem);
        } else {
            container.insertBefore(draggedItem, afterElement);
        }
    });

    container.addEventListener('drop', () => {
        const newArr = [];
        const itemDivs = container.querySelectorAll('.question-box');

        itemDivs.forEach(div => {
            const title = div.dataset.title.replace(/&quot;/g, '"');
            const originalData = currentListData.find(q => q.title === title);

            if (originalData) {
                newArr.push(originalData);
            }
        });
        selectedQuestions = newArr;
        renderSelectedQuestions();
    });
}

function getDragAfterElement(container, y) {
    const draggableElements = [...container.querySelectorAll('.question-box:not(.dragging)')];
    return draggableElements.reduce((closest, child) => {
        const box = child.getBoundingClientRect();
        const offset = y - box.top - box.height / 2;
        if (offset < 0 && offset > closest.offset) {
            return {
                offset: offset,
                element: child
            };
        } else {
            return closest;
        }
    }, {
        offset: -Infinity
    }).element;
}

document.addEventListener("DOMContentLoaded", () => {
    setSection("forms");
});