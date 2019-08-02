let reports;


function reportQuery(callback) {
    let contains = document.getElementById("placeHolder").value;
    let filter = document.getElementById("filter").value;
    let request = new XMLHttpRequest();
    request.onreadystatechange= function () {
        if(request.readyState === 4 && request.status === 200) {
            const json = JSON.parse(this.responseText);
            callback(json);
        }
    };
    request.open("GET","/reportQuery?contains="+ contains +"&filter="+ filter, true);
    request.send()
}

function displayQuery(array) {
    if(array == null) {
        noRecordFound();
        return;
    }
    var tBody = document.getElementById("lowerScreen");
    tBody.innerHTML = "";
    var body = "";
    array = array.reverse();
    array.forEach(element => {
        row = "<tr>";
        row += "<th>" +element["testSuite"] + "</th>";
        row += "<td>" +element["testFailing"] + "</td>";
        row += "<td>" +element["branch"] + "</td>";
        row += "<td>" +element["date"] + "</td>";
        row += "<td>" +element["environment"] + "</td>";
        row += "<td>" +element["rerun"] + "</td>";
        row += "<td>" +element["sameError"] + "</td>";
        row += "<td><a href=\"" +element['reportUrl'] + "\">Report</a></td>";
        row += "<td>" +element["exeHost"] + "</td>";
        row += "<td>" +element["screenDef"] + "</td>";
        row += "<td>" +element["failSince"] + "</td>";
        row += "<td>" +element["error"] + "</td>";
        row += "<td>" +element["details"] + "</td>";
        row += "<td>" +element["passRelaunch"] + "</td>";
        row += "<td>" +element["actionTaken"] + "</td>";
        row += "</tr>";
        body += row;
    });
    tBody.innerHTML = body;
}
function sendUrl() {
    document.getElementById("submit").classList.add("is-loading");
    var urlNode = document.getElementById("url");
    var url = urlNode.value;
    if(!validateUrl(url)) {
        urlNode.classList.add("is-danger");
        document.getElementById("submit").classList.remove("is-loading");
        return;
    }
    if(urlNode.classList.contains("is-danger")) {
        urlNode.classList.remove("is-danger")
    }
    var request = new XMLHttpRequest();
    request.open('POST', '/report?url=' + url, true);
    request.send();
    request.onload = function() {
        displayReport(this.response);
        document.getElementById("submit").classList.remove("is-loading");

    }
}
function sendFilledReport(id, node) {
    var formData = document.getElementById("form" + id);
    screenDefArr = [formData.getElementsByClassName("screenDef").item(0).value];
    var report = {
        "rerun" : formData.getElementsByClassName("rerun").item(0).checked.toString() ,
        "sameError": formData.getElementsByClassName("sameError").item(0).checked.toString(),
        "screenDef": screenDefArr,
        "id": formData.getElementsByClassName("id").item(0).value,
        "failSince": formData.getElementsByClassName("failSince").item(0).value,
        "errorType": formData.getElementsByClassName("errorType").item(0).value,
        "details": formData.getElementsByClassName("details").item(0).value,
        "passRelaunch": formData.getElementsByClassName("passRelaunch").item(0).checked.toString(),
        "actionTaken": formData.getElementsByClassName("actionTaken").item(0).value
    }
    var request = new XMLHttpRequest();
    request.open('POST', '/reportFill', true);
    request.send(JSON.stringify(report));

    node.parentNode.parentNode.style.display ="none";
    deleteReport(id,node);
    reports = reports.filter(function () { return true });
}
function refreshReport() {
        var request = new XMLHttpRequest();
        request.open('GET', '/reportData', true);
        request.send();
        request.onload = function() {
            displayReport(this.response);
        }
}
function displayReport(response) {
    reports = [];
    var count = 0;
    var data = JSON.parse(response);
    data.forEach(element => {
        report = {};
        report.testSuite = element.testSuite;
        report.testFailing = element.testFailing;
        report.branch = element.branch;
        report.date = element.date;
        report.environment = element.environment;
        report.reportUrl = element.reportUrl;
        report.exeHost = element.exeHost;
        report.screenDef = element.screenDef;
        report.rerun = element.rerun;
        report.sameError = element.sameError;
        report.failSince = element.failSince;
        report.error = element.error;
        report.details = element.details;
        report.passRelaunch = element.passRelaunch;
        report.actionTaken = element.actionTaken;

        report.id = count;
        report.delete = false;
        reports.push(report);
        count += 1;

    });
    displayHtml();

}


function displayHtml() {
    if (reports !== undefined) {
        var upperScreen = document.getElementById("upperScreen");
        upperScreen.innerHTML = "";

        let count = 0;

        reports.forEach(element => {
            console.log(count);
            //form
            let formNode = document.createElement("form");
            // formNode.action = "/reportFill";
            // formNode.method = "post";
            formNode.id = "form" + count;
            //upperScreen.innerHTML += "<form action=\"/reportFill\" method=\"post\">"

            //nav node
            let navNode = document.createElement("nav");
            navNode.className += "level";

            //delete Button
            let deleteButton = document.createElement("div");
            deleteButton.className += "level-item has-text-centered";
            deleteButton.innerHTML = "<a class=\"delete is-medium\" onclick=\"deleteReportRequest(" + count + ",this)\"></a>";
            navNode.appendChild(deleteButton);

            //test failling
            var testNameNode = document.createElement("div");
            testNameNode.className += "level-item has-text-centered";
            testNameNode.innerHTML = "<label onclick='queryName(displayQuery, this)' id='nameTest' class=\"label has-text-danger\">" + element.testFailing + "</label>";
            navNode.appendChild(testNameNode);

            //rerun checkbox
            var rerunNode = document.createElement("div");
            rerunNode.className += "level-item has-text-centered";
            rerunNode.innerHTML = "<label class=\"checkbox\"><b>Rerun </b><input type=\"checkbox\" class=\"rerun\"></label>";
            navNode.appendChild(rerunNode);

            //same error checkbox
            var sameErrorNode = document.createElement("div");
            sameErrorNode.className += "level-item has-text-centered";
            sameErrorNode.innerHTML = "<label class=\"checkbox\"><b>Same Error </b><input type=\"checkbox\" class=\"sameError\"></label>";
            navNode.appendChild(sameErrorNode);

            //Screendef option
            var screenDefNode = document.createElement("div");
            screenDefNode.className += "level-item has-text-centered";

            var option = "<div class=\"field\"><div class=\"control\"><div class=\"select is-primary\"><select class=\"screenDef\"><option>ScreenDef</option>";
            option += "<option value='None'>None</option>";
            if(element.screenDef != null) {
                element.screenDef.forEach(item => {
                    option += "<option value=" + item + ">" + item + "</option>";
                });
            };
            option += "</select></div></div></div>";
            screenDefNode.innerHTML = option;
            navNode.appendChild(screenDefNode);

            //fail since
            var failSinceNode = document.createElement("div");
            failSinceNode.className += "level-item has-text-centered";
            failSinceNode.innerHTML = "<div class=\"field\"><input type='hidden' class='id' value='"+count+"'><div class=\"control\"><input class=\"input is-primary failSince\" type=\"text\" placeholder=\"Fail since\"></div></div>";
            navNode.appendChild(failSinceNode);

            //error type
            var errorNode = document.createElement("div");
            errorNode.className += "level-item has-text-centered";
            errorNode.innerHTML = "<div class=\"field\">"
                + "<div class=\"control\">"
                + "<div class=\"select is-primary\">"
                + "<select class=\"errorType\">"
                + "<option>Error</option>"
                + "<option value=\"Instability\">Instability</option>"
                + "<option value=\"Assert\">Assert</option>"
                + "<option value=\"Technical\">Technical</option>"
                + "<option value=\"Flow\">Flow</option>"
                + "<option value=\"Administration\">Administration</option>"
                + "<option value=\"WrongVersioning\">Wrong Versioning</option>"
                + "</select>"
                + "</div>"
                + "</div>"
                + "</div>";
            navNode.appendChild(errorNode);

            //details node
            var detailsNode = document.createElement("div");
            detailsNode.className += "level-item has-text-centered";
            detailsNode.innerHTML = "<div class=\"field\">"
                + "<div class=\"control\">"
                + "<input class=\"input is-primary details\" type=\"text\" placeholder=\"Details\">"
                + "</div>"
                + "</div>";
            navNode.appendChild(detailsNode);

            //pass relaunch
            var passNode = document.createElement("div");
            passNode.className += "level-item has-text-centered";
            passNode.innerHTML = "<label class=\"checkbox\"><b>Pass relaunch </b><input type=\"checkbox\" class=\"passRelaunch\"></label>";
            navNode.appendChild(passNode);

            //action taken
            var actionNode = document.createElement("div");
            actionNode.className += "level-item has-text-centered";
            actionNode.innerHTML = "<div class=\"field\">"
                + "<div class=\"control\">"
                + "<input class=\"input is-primary actionTaken\" type=\"text\" placeholder=\"Action taken\">"
                + "</div>"
                + "</div>";
            navNode.appendChild(actionNode);

            //submit
            var submitNode = document.createElement("div");
            submitNode.className += "level-item has-text-centered";
            submitNode.innerHTML = "<button  class=\"button is-primary\" type='button' onclick=\"sendFilledReport(" +count+",this)\" >Submit</button>";
            navNode.appendChild(submitNode);


            //line break
            var linebreak = document.createElement("hr");

            formNode.appendChild(navNode);
            formNode.appendChild(linebreak);

            //formNode.appendChild(navNode);
            upperScreen.appendChild(formNode);
            count += 1;
        });
    }
}

function deleteReport(id, node) {
    node.parentNode.parentNode.parentNode.style.display ="none";
    let indexId
    for (i = 0; i < reports.length; i++)  {
        if(reports[i].id === id) {
            indexId = i;
        }
    }
    reportDel = reports[indexId];
    delete reports[indexId];

    return reportDel;
}

function deleteReportRequest(id,node) {
    reportDel = deleteReport(id,node);

    var request = new XMLHttpRequest();
    request.open('POST', '/reportDelete', true);
    request.send(JSON.stringify(reportDel));
    reports = reports.filter(function () { return true });
}

function queryName(callback, node) {
    let filter = "Test failling";
    let contains = node.innerText;
    console.log(contains);
    let request = new XMLHttpRequest();
    request.onreadystatechange= function () {
        if(request.readyState === 4 && request.status === 200) {
            const json = JSON.parse(this.responseText);
            callback(json);
        }
    };
    request.open("GET","/reportQuery?contains="+ contains +"&filter=" + filter, true);
    request.send()
}
function validateUrl(url)  {
    var patternBDI = new RegExp('.*BDI.*');
    var patternUrl = new RegExp('^(https:\/\/stha38e56:444\/view\/Billing\/view\/3\.Scheduled%20Tests).*(HTML-Report\/)$');
    return patternUrl.test(url) && !patternBDI.test(url);
}

function noRecordFound() {
    let modal = document.getElementById("noRecordModal");
    modal.classList.add("is-active");
}
function closeModal() {
    let modal = document.getElementById("noRecordModal");
    modal.classList.remove("is-active");
}


