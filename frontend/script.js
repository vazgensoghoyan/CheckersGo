// ====== Переменные ======
let playerId = null;
let playerColor = null;
let boardData = [];
let currentTurn = null;
let selectedCell = null;

// ====== Элементы DOM ======
const connectBtn = document.getElementById("connect-btn");
const colorSelect = document.getElementById("color-select");
const connectStatus = document.getElementById("connect-status");
const playerInfo = document.getElementById("player-info");
const playerColorElem = document.getElementById("player-color");
const currentTurnElem = document.getElementById("current-turn");
const boardSection = document.getElementById("board-section");
const boardTable = document.getElementById("board");
const winnerText = document.getElementById("winner-text");
const gameStatus = document.getElementById("game-status");

// ====== Функции ======

// Подключение игрока
async function connectPlayer() {
    const color = colorSelect.value;
    try {
        const response = await fetch("/connect", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ isWhite: color === "white" })
        });
        const data = await response.json();

        if (data.error) {
            connectStatus.textContent = "Ошибка: " + data.error;
            return;
        }

        // Сохраняем данные игрока
        playerId = data.id;
        playerColor = data.color;
        currentTurn = data.turn;

        // Показываем интерфейс
        connectStatus.textContent = "Подключение успешно!";
        playerInfo.style.display = "block";
        boardSection.style.display = "block";
        gameStatus.style.display = "block";
        playerColorElem.textContent = playerColor;
        currentTurnElem.textContent = currentTurn;

        // Загружаем доску
        await loadBoard();

        // Запускаем авто-обновление
        setInterval(loadBoard, 1000);

    } catch (err) {
        connectStatus.textContent = "Ошибка подключения";
        console.error(err);
    }
}

// Загрузка доски с сервера
async function loadBoard() {
    try {
        const response = await fetch("/board");
        const data = await response.json();

        boardData = data.board;
        currentTurn = data.turn;
        currentTurnElem.textContent = currentTurn;

        if (data.winner) {
            winnerText.textContent = "Победитель: " + data.winner;
        } else {
            winnerText.textContent = "";
        }

        renderBoard();
    } catch (err) {
        console.error("Ошибка загрузки доски", err);
    }
}

// Отрисовка доски
function renderBoard() {
    boardTable.innerHTML = "";

    for (let i = 0; i < 8; i++) {
        const row = document.createElement("tr");
        for (let j = 0; j < 8; j++) {
            const cell = document.createElement("td");
            cell.dataset.row = i;
            cell.dataset.col = j;

            // Добавляем шашку
            const cellValue = boardData[i][j];
            if (cellValue === 1) {
                const piece = document.createElement("div");
                piece.className = "piece-white";
                cell.appendChild(piece);
            } else if (cellValue === -1) {
                const piece = document.createElement("div");
                piece.className = "piece-black";
                cell.appendChild(piece);
            }

            // Клик по клетке
            cell.addEventListener("click", () => handleCellClick(i, j));
            row.appendChild(cell);
        }
        boardTable.appendChild(row);
    }
}

// Обработка клика по клетке
function handleCellClick(row, col) {
    if (currentTurn !== playerColor) return; // не наш ход

    const cellValue = boardData[row][col];

    // Если нажали на свою шашку — выбираем
    if ((playerColor === "white" && cellValue === 1) ||
        (playerColor === "black" && cellValue === -1)) {
        selectedCell = { row, col };
        highlightSelected();
        return;
    }

    // Если выбрана шашка и клик на другую клетку — делаем ход
    if (selectedCell) {
        makeMove(selectedCell, { row, col });
        selectedCell = null;
        renderBoard();
    }
}

// Подсветка выбранной шашки
function highlightSelected() {
    renderBoard();
    if (!selectedCell) return;
    const cell = boardTable.rows[selectedCell.row].cells[selectedCell.col];
    cell.classList.add("selected");
}

// Отправка хода на сервер
async function makeMove(from, to) {
    try {
        const response = await fetch("/board/move", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                id: playerId,
                from: [from.row, from.col],
                to: [to.row, to.col]
            })
        });
        const data = await response.json();

        if (data.error) {
            alert("Ошибка: " + data.error);
            return;
        }

        boardData = data.board;
        currentTurn = data.turn;
        currentTurnElem.textContent = currentTurn;

        if (data.winner) {
            winnerText.textContent = "Победитель: " + data.winner;
        }

        renderBoard();
    } catch (err) {
        console.error("Ошибка хода", err);
    }
}

// ====== События ======
connectBtn.addEventListener("click", connectPlayer);
