// logic.js — упрощённая версия для двух игроков на одной странице

const API_BASE = "http://localhost:8080"; // если сервер на другом хосте, например "http://localhost:8080"

const boardEl = document.getElementById("board");
const statusEl = document.getElementById("status");
const msgEl = document.getElementById("serverMsg");

let boardState = null;

let player1 = { id: null, color: null };
let player2 = { id: null, color: null };

// --- DOM helpers ---
function $(id){ return document.getElementById(id); }
function idxToCoord(r,c){ return String.fromCharCode(97 + c) + (8 - r); }
function setMsg(t){ msgEl.textContent = t; }
function getField(obj, ...names){ for(const n of names) if(obj && obj[n]!==undefined) return obj[n]; return undefined; }

// --- API ---
async function requestJSON(path, options={}) {
  const url = (API_BASE||"")+path;
  const res = await fetch(url, options);
  const text = await res.text();
  let json;
  try { json = text ? JSON.parse(text) : {}; } catch(e){ throw new Error("Invalid JSON"); }
  if(!res.ok){ const msg=getField(json,"error","Error")||`HTTP ${res.status}`; throw new Error(msg); }
  return json;
}

async function apiJoin(name){ return requestJSON("/join",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({name})});}
async function apiState(){ return requestJSON("/state"); }
async function apiMove(player_id, from, to){ return requestJSON("/move",{method:"POST",headers:{"Content-Type":"application/json"},body:JSON.stringify({player_id, from, to})});}
async function apiReset(){ return requestJSON("/reset",{method:"POST"});}

// --- Board ---
function buildBoard(){
  boardEl.innerHTML="";
  for(let r=0;r<8;r++){
    for(let c=0;c<8;c++){
      const sq=document.createElement("div");
      sq.className="square "+((r+c)%2===0?"light":"dark");
      sq.textContent=idxToCoord(r,c);
      boardEl.appendChild(sq);
    }
  }
}

function normalizeFigure(f){ return f?{is_none:getField(f,"is_none","IsNone"),is_white:getField(f,"is_white","IsWhite"),is_king:getField(f,"is_king","IsKing")}: {is_none:true,is_white:false,is_king:false}; }

function renderBoard(state){
  const board = getField(state,"board","Board");
  const isWhiteTurn = getField(state,"isWhiteTurn","IsWhiteTurn","is_white_turn")||false;
  if(!board||!Array.isArray(board)){ setMsg("Некорректный ответ /state"); return; }
  boardState = board;
  const squares = boardEl.children;
  for(let r=0;r<8;r++){
    for(let c=0;c<8;c++){
      const idx=r*8+c;
      const sq=squares[idx];
      sq.innerHTML=idxToCoord(r,c);
      const f=normalizeFigure(board[r][c]);
      if(!f.is_none){
        const p=document.createElement("div");
        p.className="piece "+(f.is_white?"white":"black")+(f.is_king?" king":"");
        p.textContent=f.is_king? (f.is_white?"♔":"♚"):"";
        sq.appendChild(p);
      }
    }
  }
  statusEl.textContent=`Ход: ${isWhiteTurn?"Белые":"Черные"}`;
}

// --- Player actions ---
async function joinPlayer(player, name, infoEl){
  const res = await apiJoin(name);
  player.id = getField(res,"player_id","PlayerID");
  player.color = getField(res,"color","Color");
  infoEl.textContent = `ID: ${player.id}, Цвет: ${player.color}`;
  setMsg(`Игрок ${name} присоединился как ${player.color}`);
  await fetchAndRender();
}

async function movePlayer(player, from, to){
  if(!player.id){ setMsg("Сначала join"); return; }
  try{
    const res = await apiMove(player.id, from, to);
    setMsg(res.message || "Ход выполнен");
    await fetchAndRender();
  }catch(e){ setMsg("Ошибка /move: "+e.message); }
}

// --- Fetch board ---
async function fetchAndRender(){
  try{ const st = await apiState(); renderBoard(st); } catch(e){ setMsg("Ошибка /state: "+e.message); }
}

// --- Buttons ---
$("join1").onclick = ()=>joinPlayer(player1, $("name1").value.trim(), $("info1"));
$("move1").onclick = ()=>movePlayer(player1, $("from1").value.trim(), $("to1").value.trim());

$("join2").onclick = ()=>joinPlayer(player2, $("name2").value.trim(), $("info2"));
$("move2").onclick = ()=>movePlayer(player2, $("from2").value.trim(), $("to2").value.trim());

$("refreshBtn").onclick = fetchAndRender;
$("resetBtn").onclick = async ()=>{
  if(!confirm("Сбросить игру?")) return;
  await apiReset();
  player1={id:null,color:null}; player2={id:null,color:null};
  $("info1").textContent="—"; $("info2").textContent="—";
  setMsg("Игра сброшена");
  fetchAndRender();
};

// --- Init ---
buildBoard();
fetchAndRender();
