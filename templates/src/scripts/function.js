const server = '/';

$(function(){
    $('a[href^="#"]').on('click', function(event) {
        // отменяем стандартное действие
        event.preventDefault();

        var sc = $(this).attr("href"),
            dn = $(sc).offset().top;
        /*
        * sc - в переменную заносим информацию о том, к какому блоку надо перейти
        * dn - определяем положение блока на странице
        */

        $('html, body').animate({scrollTop: dn}, 1800);

        /*
        * 1000 скорость перехода в миллисекундах
        */
    });
});

async function ListUpdate(value) {
}

async function selectCub(value) {
    console.log(value);
    let el = document.getElementById("Cub_" + value);
    document.querySelectorAll('#TypeBlocks .block').forEach(n => n.classList.remove('active'));
    document.getElementById("RadiusBlocksFront").innerHTML = '<p style="font-size: .6em; padding: 10px 0;">\n                                Сначала выберите тип и кубатуру\n                            </p>';
    document.getElementById("RadiusBlocksBack").innerHTML = '';
    document.querySelectorAll('#CubBlocks .block').forEach(n => n.classList.remove('active'));

    el.classList.add('active');
    setCookie('cub', value, 'Tue, 19 Jan 2038 03:14:07 GMT', '/');

    deleteCookie('type');
    deleteCookie('radius_front');
    deleteCookie('radius_back');
}

async function selectType(value) {
    setCookie('type', value, 'Tue, 19 Jan 2038 03:14:07 GMT', '/');
    document.querySelectorAll('#TypeBlocks .block').forEach(n => n.classList.remove('active'));
    document.querySelectorAll('#RadiusBlocksFront').forEach(n => n.classList.remove('full_radius'));
    document.querySelectorAll('#RadiusBlocksBack').forEach(n => n.classList.remove('full_radius'));

    document.getElementById("Type_" + value).classList.add('active');

    document.getElementById("RadiusBlocksFront").innerHTML = '';
    document.getElementById("RadiusBlocksBack").innerHTML = '';
    document.getElementById("ProductPrice").innerHTML = '';

    if(value === 2){
        document.getElementById("RadiusBlocksFront").classList.add('full_radius');
    }else if(value === 3){
        document.getElementById("RadiusBlocksBack").classList.add('full_radius');
    }

    deleteCookie('radius_front');
    deleteCookie('radius_back');
    deleteCookie('spike');

    await selectRadius();
}

async function selectRadius() {
    let requestURL = server + 'tires/cub/' + getCookie("cub") + '/type/' + getCookie("type");
    let request = await fetch(requestURL, {
        method: "GET"
    });
    let TiresRadius = await request.json();

    document.getElementById("RadiusBlocksFront").innerHTML = '';
    document.getElementById("RadiusBlocksBack").innerHTML = '';
        let clickCount = 0;
        let type = getCookie("type");
        let arrRad;

        for (k in TiresRadius){
            arrRad = TiresRadius[k]["RadiusFront"];
        }
        let unique = [ ...new Set([arrRad])];
        let i = 0;
        let j = 0;
        let spikes;
        for (key in TiresRadius){
            let div = document.createElement('div');
            div.className = "block";
            let back = document.createElement('div');
            back.className = "block";
            let column = document.createElement('div');
            column.className = "col-lg-3 col-md-3 col-sm-4 col-6";

            if(type === "1"){
                column.className = "col-lg-6 col-md-6 col-sm-6 col-12";
                console.log("key:" + key);
                //if(TiresRadius[key]["RadiusBack"] && TiresRadius[key]["RadiusBack"] !== '' && TiresRadius[key]["RadiusBack"] !== ' '){

                //}

                if(unique[i]) {
                    let col = document.createElement('div');
                    col.className = "col-lg-6 col-md-6 col-sm-6 col-12";

                    spikes = TiresRadius[key]["Spike"] === 0 ? "Передняя" : TiresRadius[key]["Spike"] + " шипов";
                    div.id = "RadiusFront_" + unique[i];
                    div.innerHTML = "<h4>" + unique[i] + " R</h4> <p>" + spikes + "</p>";
                    document.getElementById("RadiusBlocksFront").appendChild(col);
                    document.getElementById("RadiusBlocksFront").children[i].appendChild(div);
                    i++;
                    div.onclick = function(){
                        clickCount++;
                        document.querySelectorAll('#RadiusBlocksFront .block').forEach(n => n.classList.remove('active'));
                        div.classList.add('active');
                        let divId = div.id;
                        let arr = divId.split('_');
                        setCookie('radius_front', arr[1], 'Tue, 19 Jan 2038 03:14:07 GMT', '/');

                        if (clickCount === 2){ selectPrice(); }
                    };
                }else{
                    document.getElementById("RadiusBlocksFront").appendChild(column);
                }

                spikes = TiresRadius[key]["Spike"] === 0 ? "Задняя" : TiresRadius[key]["Spike"] + " шипов";
                back.id = "RadiusBack_" + TiresRadius[key]["RadiusBack"];
                back.innerHTML = "<h4>" + TiresRadius[key]["RadiusBack"] + " R</h4> <p>" + spikes + "</p>";
                //document.getElementById("RadiusBlocksBack").appendChild(back);
                document.getElementById("RadiusBlocksBack").appendChild(column);
                document.getElementById("RadiusBlocksBack").children[j].appendChild(back);
                console.log("j:" + j)
                j++;
                back.onclick = function(){
                    clickCount++;
                    document.querySelectorAll('#RadiusBlocksBack .block').forEach(n => n.classList.remove('active'));
                    back.classList.add('active');
                    let divId = back.id;
                    let arr = divId.split('_');
                    setCookie('radius_back', arr[1], 'Tue, 19 Jan 2038 03:14:07 GMT', '/');

                    if (clickCount === 2){ selectPrice(); }
                }
            }else if(type === "2"){
                spikes = TiresRadius[key]["Spike"] === 0 ? "Передняя" : TiresRadius[key]["Spike"] + " шипов";
                div.id = "RadiusFront_" + TiresRadius[key]["RadiusFront"];
                div.innerHTML = "<h4>" + TiresRadius[key]["RadiusFront"] + " R</h4> <p>" + spikes + "</p>";
                document.getElementById("RadiusBlocksFront").appendChild(column)
                document.getElementById("RadiusBlocksFront").children[j].appendChild(div);
                div.onclick = function(){
                    document.querySelectorAll('#RadiusBlocksBack .block').forEach(n => n.classList.remove('active'));
                    div.classList.add('active');
                    let divId = div.id;
                    let arr = divId.split('_');
                    setCookie('radius_front', arr[1], 'Tue, 19 Jan 2038 03:14:07 GMT', '/');
                    selectPrice();
                };
                j++;
            }else {
                spikes = TiresRadius[key]["Spike"] === 0 ? "Задняя" : TiresRadius[key]["Spike"] + " шипов";
                if(TiresRadius[key]["Spike"] === 0){
                    div.id = "RadiusBack_" + TiresRadius[key]["RadiusBack"];
                }else{
                    div.id = "RadiusBack_" + TiresRadius[key]["RadiusBack"] + "_Spike_" + TiresRadius[key]["Spike"];
                }
                div.innerHTML = "<h4>" + TiresRadius[key]["RadiusBack"] + " R</h4> <p>" + spikes + "</p>";
                document.getElementById("RadiusBlocksBack").appendChild(column)
                document.getElementById("RadiusBlocksBack").children[j].appendChild(div);

                div.onclick = function(){
                    document.querySelectorAll('#RadiusBlocksBack .block').forEach(n => n.classList.remove('active'));
                    div.classList.add('active');
                    let divId = div.id;
                    let arr = divId.split('_');
                    setCookie('radius_back', arr[1], 'Tue, 19 Jan 2038 03:14:07 GMT', '/');
                    setCookie('spike', arr[3], 'Tue, 19 Jan 2038 03:14:07 GMT', '/');
                    selectPrice();
                };
                j++;
            }
        }//endFor
}

async function selectPrice() {
    let type = getCookie("type");
    let requestURL;
    document.getElementById("ProductPrice").innerHTML = '';

    if(type === "1"){
        requestURL = server + 'tires/cub/' + getCookie("cub") + '/type/' +
            getCookie("type") + '/rFront/' + getCookie("radius_front") + '/rBack/' + getCookie("radius_back");
    }else if(type === "2"){
        requestURL = server + 'tires/cub/' + getCookie("cub") + '/type/' +
            getCookie("type") + '/rFront/' + getCookie("radius_front");
    }else{
        let spike = getCookie("spike") ? getCookie("spike") : 0;
        requestURL = server + 'tires/cub/' + getCookie("cub") + '/type/' +
            getCookie("type") + '/rBack/' + getCookie("radius_back") + '/spike/' + spike;
    }

    let request = await fetch(requestURL, {
        method: "GET"
    });
    let TiresPrice = await request.json();

    let div = document.createElement('div');
    div.className = "price";
    div.innerHTML = "<h4>Цена: " + TiresPrice['Price'] + "</h4> <button onclick=\"sendEmail(" + TiresPrice['Id'] + ", 'tires')\">Купить</button>";
    document.getElementById("ProductPrice").appendChild(div);
}

async function sendEmail(ProductId, ProductName){
    let requestURL = server + 'product/' + ProductName + '/id/' + ProductId;

    let request = await fetch(requestURL, {
        method: "GET"
    });
    let orderId = await request.json();

    document.getElementById("popup").style.display = "flex";
    document.getElementById("popupForm").action = "/orderId/" + orderId;
    //onclick="window.location.href = '/';"
}

function hidePopup() {
    document.getElementById("popup").style.display = "none";
}