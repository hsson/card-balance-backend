package balance

import "testing"

const exampleLoginPage = `
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
    <head>
        <title>
            Logga in
        </title>
        <link rel="icon" href="styles/Images/favicon.ico" type="image/x-icon" />
        <link rel="stylesheet" type="text/css" href="Content/bootstrap-theme.css" />
        <link rel="stylesheet" type="text/css" href="Content/bootstrap.css" />
        <script src="jquery/external/jquery/jquery.js"></script>
        <script src="Scripts/bootstrap.js"></script>
        <script src="Scripts/bootbox.js"></script>
        <link rel="stylesheet" type="text/css" href="styles/Netaxept.css" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <!-- Begin Cookie Consent plugin by Silktide - http://silktide.com/cookieconsent -->
        <script type="text/javascript">
            window.cookieconsent_options = { "message": "Denna webbplats använder cookies för att säkerställa att du får den bästa upplevelsen på vår hemsida. Genom att fortsätta använda vår hemsida godkänner du vår cookie policy.", "dismiss": "Jag förstår!", "learnMore": "Mer information", "link": null, "theme": "light-bottom" };
        </script>
        <script type="text/javascript" src="Scripts/cookieconsent.min.js"></script>
        <!-- End Cookie Consent plugin -->
    </head>
    <body>
        <div class="wrapper">
            <form name="frmLogin" method="post" action="default.aspx" id="frmLogin">
                <div>
                    <input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="/wEPDwUKMTEzNTI2MjQyMg9kFgQCAw9kFgoCAQ9kFgICAQ8PFgIeBFRleHQFClBUTSBLb5J0bnJkZAICDxYCHgdWaXNpYmxlaGQCAkadl88BaGQCBA8WAh8BIdjwaQ8WAh8BaBYCAgEPEGRkFgBkAgUPDxYCHwAFCShkZXNrdG9wKWRkZMD6W4zuUf8/+Dp04d7fIVjURaNB" />
                </div>
                <div>
                    <input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="CX0A0534" />
                    <input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="/wEWBAKNp6aXDwLi0uqnCgKF59rWBAK94jAbChDpRor7Q17FkzZGQUmFZy1hHJ2a" />
                </div>
                <input type="hidden" name="hiddenIsMobile" id="hiddenIsMobile" value="" />
                <div class="container">
                    <div class="title">
                        <img alt="PTMweb logga" src="styles/images/PTMWeb_utan_skugga.png" class="TitleIcon" />
                        <div id="titleText" class="TitleText">Kortladdning</div>
                    </div>
                </div>
                <div class="container">
                    <br />
                    <div id="CardNumberSection">
                        <div class="row">
                            <div class="col-md-3">
                                <span id="lblCardnum">PTM Kortnr</span>:
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-md-3">
                                <input name="txtCardNumber" type="text" id="txtCardNumber" Class="form-control" />
                                <b class="Default_ErrorMessageText">
                                <br />
                                <span id="lblCardMissing"></span>
                                <input type="hidden" name="SavedCardNumber" id="SavedCardNumber" />
                                </b>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-md-3">
                                <input type="submit" name="btnNext" value="Nästa" id="btnNext" Class="btn btn-primary" />
                            </div>
                        </div>
                    </div>
                    <br />
                    <br />
                    <a href="docs/manual.pdf" id="manualLink">klicka här för att se manualen.</a>
                    <div class="footer">
                        <div class="row">
                            <div class="col-md-9">
                                <p>Chalmers Konferens och Restauranger, Sven Hultins gata 4, 412 58 Göteborg.</p>
                            </div>
                            <div class="col-md-3" style="float: right">
                                <img src="styles/images/Atronic-logo-RGB-tagline.png" alt="PTMweb logga" aria-autocomplete="none" />
                            </div>
                        </div>
                    </div>
                </div>
            </form>
        </div>
        <small style="float: right; visibility: hidden; font-size: 8px;">
        <span id="lblIsMobile">(desktop)</span></small>
        <script type="text/javascript">
            var r = window.devicePixelRatio;
            if (typeof r == "undefined")
                r = 1;
            
            var size = screen.width * r;
            if (screen.height * r > size)
                size = screen.height * r;
            
            document.getElementById("hiddenIsMobile").value = (size <= 1000 ? "mobile" : "desktop");
        </script>
        <script>
            var d = new Date();
            document.getElementById("manualLink").href =
              "docs/manual.pdf?ver=" + d.getTime();
        </script>
    </body>
</html>
`

const exampleCardDetailsPage = `
<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml">
    <head>
        <title>
            Order
        </title>
        <meta content="width=device-width, initial-scale=1" name="viewport" />
        <link href="styles/Images/favicon.ico" rel="icon" type="image/x-icon" />
        <link href="Content/bootstrap-theme.css" rel="stylesheet" type="text/css" />
        <link href="Content/bootstrap.css" rel="stylesheet" type="text/css" />
        <link href="styles/Netaxept.css" rel="stylesheet" type="text/css" />
        <script src="jquery/external/jquery/jquery.js"></script>
        <link href="jquery/jquery-ui.css" rel="stylesheet" />
        <script src="jquery/jquery-ui.js"></script>
        <link href="jquery/jquery-ui.structure.css" rel="stylesheet" />
        <link href="jquery/jquery-ui.theme.css" rel="stylesheet" />
        <script src="Scripts/bootstrap.js"></script>
        <script src="Scripts/bootbox.js"></script>
        <script>
            $(function () {
                $("#txtStatementStartDate").datepicker({ dateFormat: 'yy-mm-dd' }).val();
            });
            
            $(function () {
                $("#txtStatementStopDate").datepicker({ dateFormat: 'yy-mm-dd' }).val();
            });
        </script>
    </head>
    <body>
        <div class="wrapper">
            <form name="frmOrder" method="post" action="CardLoad_Order.aspx" id="frmOrder">
                <div>
                    <input type="hidden" name="__EVENTTARGET" id="__EVENTTARGET" value="" />
                    <input type="hidden" name="__EVENTARGUMENT" id="__EVENTARGUMENT" value="" />
                    <input type="hidden" name="__VIEWSTATE" id="__VIEWSTATE" value="/wEPDwUKLTQyMDkyMTczMA9kFgICAw9kFjICAw8PFgIeBFRleHQFFEFsZXhhbmRlciBIw6VrYW5zc29uZGQCBQ8PFgIfAAUVYWwzeGFuZDNyLmhAZ21haWwuY29tZGQCBw8PFgIfAAUQMzgxOTI4NTE4NTEzODk0MmRkAgkPDxYCHwAFBTYzLDc4ZGQCDA8WBB4IZGlzYWJsZWRkHgdWaXNpYmxlaGQCDQ8PFgQeCUZvcmVDb2xvcgojHgRfIVNCAgRkZAIODw8WBB8DCiMfBAIEZGQCDw8PFgYfAwpOHgdFbmFibGVkaB8EAgRkZAIQDw8WBB8DCiMfBAIEZGQCEQ8PFgYfAwojHwQCBB8CaGRkAhIPDxYEHwMKIx8EAgRkZAITDw8WAh8CaGRkAhQPDxYCHwJoZGQCFQ8WAh8CaBYCAgMPEGQPFgpmAgECAgIDAgQCBQIGAgcCCAIJFgoQBQMyMDAFAzIwMGcQBQM0MDAFAzQwMGcQBQM2MDAFAzYwMGcQBQM4MDAFAzgwMGcQBQQxMDAwBQQxMDAwZxAFBDEyMDAFBDEyMDBnEAUEMTQwMAUEMTQwMGcQBQQxNjAwBQQxNjAwZxAFBDE4MDAFBAapwgBnEAUEMjAwMAUEMjAwMGcWAWZkAhYPFgIfAmgWAgIDDxBkDxYEZgIBAgICAxYEEAUpTsOkciBrb3J0dsOkcmRlIHVuZGVyc3RpZ2VyIGzDpGdzdGEgbml2w6UFATFnEAUSRsO2cnN0YSBpIG3DpW5hZGVuBQEyZxAFEFNpc3RhIGkgbcOlbmFkZW4FATNnEAUXU3BlY2lmaWsgZGFnIGkgbcOlbmFkZW4FATRnFgFmZAIXDxYCHwJoFgICAw8QZA8WCmYCAQICAgMCBAIFAgYCBwIIAgkWChAFAzEwMAUDMTAwZxAFAzIwMAUDMjAwZxAFAzMwMAUDMzAwZxAFAzQwMAUDNDAwZxAFAzUwMAUDNTAwZxAsjwhGMAUDNjAwZxAFAzcwMAUDNzAwZxAFAzgwMAUDODAwZxAFAzkwMAUDOTAwZxAFBDEwMDAFBDEwMDBnFgFmZAIYDxYCHwJoFgICAw8QZA8WBWYCAQICAgMCBBYFEAUDMjAwBQMyMDBnEAUDNDAwBQM0MDBnEAUDNjAwBQM2MDBnEAUDODAwBQM4MDBnEAUEMTAwMAUEMTAwMGcWAWZkAhkPFgIfAmgWAgIDDxBkDxYcZgIBAgICAwIEAgUCBgIHAggCCQIKAgsCDAINAga7GwIQAhECEgITAhQCFQIWAhcCGAIZAhoCGxYcEAUBMQUBMWcQBQEyBQEyZxAFATMFATNnEAUBNAUBNGcQBQE1BQE1ZxAFATYFATZnEAUBNwUBN2cQBQE4BQE4ZxAFATkFATlnEAUCMTAFAjEwZxAFAjExBQIxMWcQBQIxMgUCMTJnEAUCMTMFAjEzZxAFAjE0BQIxNGcQBQIxNQUCMTVnEAUCMTYFAjE2ZxAFAjE3BQIxN2cQBQIxOAUCMThnEAUCMTkFAjE5ZxAFAjIwBQIyMGcQBQIyMQUCMjFnEAUCMjIFAjIyZxAFAjIzBQIyM2cQBQIyNAUCMjRnEAUCMjUFAjI1ZxAFAjI2BQIyNmcQBQIyNwUCMjdnEAUCMjgFAjI4ZxYBZmQCGg8WAh8CaGQCGw8WAh8CaGQCHA1WAh8CaGQCHQ8WAh8CaGQCHg8WAh8CaGQCHw8WAh8CaBYGAgMPEA8WAh4HQ2hlY2tlZGhkZGRkAgUPEA8WAh8GaGRkZGQCBw8PFgQfAwojHwQCBGRkAiAPFgIfAmhkZFHjeUU0wsiaoMg9A4kTCbX/Frr2" />
                </div>
                <script type="text/javascript">
                    //<![CDATA[
                    var theForm = document.forms['frmOrder'];
                    if (!theForm) {
                        theForm = document.frmOrder;
                    }
                    function __doPostBack(eventTarget, eventArgument) {
                        if (!theForm.onsubmit || (theForm.onsubmit() != false)) {
                            theForm.__EVENTTARGET.value = eventTarget;
                            theForm.__EVENTARGUMENT.value = eventArgument;
                            theForm.submit();
                        }
                    }
                    //]]>
                </script>
                <div>
                    <input type="hidden" name="__VIEWSTATEGENERATOR" id="__VIEWSTATEGENERATOR" value="2A22720X" />
                    <input type="hidden" name="__EVENTVALIDATION" id="__EVENTVALIDATION" value="/wEWCALX/YavCXKm2A3dCQK8+L7gDwL1zrrcAgKd1au+CgLku4/3CALVhafwAgK21OewCWzUHQUp7BMcoNSry6McI3bCwnLc" />
                </div>
                <div class="container">
                    <div class="staticInformation">
                        <input type="button" name="btnLogout" value="Logga ut" onclick="javascript:__doPostBack('btnLogout','')" id="btnLogout" Class="btn btn-danger" />
                    </div>
                    <div class="title">
                        <img alt="PTMweb logga" src="styles/images/PTMWeb_utan_skugga.png" class="TitleIcon" />
                        <div id="titleText" class="TitleText">Kortladdning</div>
                    </div>
                    <div class="staticInformation">
                        <table class="table table-bordered table-striped">
                            <tr>
                                <td>
                                    <span id="Label2">Namn:</span>
                                </td>
                                <td>
                                    <b>
                                    <span id="txtPTMCardName">John Doe</span>
                                    </b>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <span id="lblEmailLabel">E-postadress:</span>
                                </td>
                                <td>
                                    <b>
                                    <span id="lblEmail">john@example.com</span>
                                    </b>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <span id="lblCardNumber">Kortnummer:</span>
                                </td>
                                <td>
                                    <b>
                                    <span id="txtPTMCardNumber">2222333344445555</span>
                                    </b>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <span id="lblPTMCardValuePrefix">Kortvärde:</span>
                                </td>
                                <td>
                                    <b>
                                    <span id="txtPTMCardValue">69,42</span>
                                    <span id="lblPTMCardValueUnit">kr</span>
                                    </b>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <span id="Label9">Periodiskt:</span>
                                </td>
                                <td>
                                    <b>Ej aktivt just nu.</b>
                                </td>
                            </tr>
                        </table>
                    </div>
                </div>
                <div class="container">
                    <div class="CardLoad_Order_buttons">
                        <input type="submit" name="btnOrderSingle" value="Ladda en gång" id="btnOrderSingle" Class="btn btn-default center-block" style="color:Black;" />
                        <input type="submit" name="btnOrderMulti" value="Ladda periodiskt" id="btnOrderMulti" Class="btn btn-default center-block" style="color:Black;" />
                        <input type="submit" name="btnOrderMultiStop" value="Avsluta periodiskt" id="btnOrderMultiStop" disabled="disabled" Class="btn btn-default center-block" style="color:Gray;" />
                        <input type="submit" name="btnEndAutomation" value="Hantera laddning" id="btnEndAutomation" Class="btn btn-default center-block" style="color:Black;" />
                        <input type="submit" name="btnAccountStatements" value="Kontoutdrag" id="btnAccountStatements" Class="btn btn-default center-block" style="color:Black;" />
                    </div>
                </div>
                <div class="container">
                    <!-- Modal -->
                    <div id="qrModal" class="modal fade" role="dialog">
                        <div class="modal-dialog modal-sm">
                            <!-- Modal content-->
                            <div class="modal-content">
                                <div class="modal-header">
                                    <button type="button" class="close" data-dismiss="modal">&times;</button>
                                    <h4 class="modal-title">Betala med QR-kod</h4>
                                </div>
                                <div class="modal-body">
                                    <div id="qrCodePlaceholder" class="row">
                                        <div class="col-md-12" style="text-align: center; vertical-align: middle;">
                                            <img src="myqr.qr?u=qrCode.72.69.51.49.99.105.112.104.108.115.120.118.99.69.89.111.119.97.43.79.56.82.104.114.113.111.88.81.66.120.119.82" id="qrCodeHolder" alt="This is where your QR-code is supposed to be." title="" style="width: 100%; height: 50%;" />
                                        </div>
                                    </div>
                                </div>
                                <div class="modal-footer">
                                    <input type="submit" name="qrRefresh" value="Uppdatera" id="qrRefresh" Class="btn btn-primary" />
                                    <button type="button" class="btn btn-default" data-dismiss="modal">Stäng</button>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="footer">
                        <div class="row">
                            <div class="col-md-9">
                                <p>Chalmers Konferens och Restauranger, Sven Hultins gata 4, 412 58 Göteborg.</p>
                            </div>
                            <div class="col-md-3" style="float: right">
                                <img src="styles/images/Atronic-logo-RGB-tagline.png" alt="PTMweb logga" aria-autocomplete="none" />
                            </div>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    </body>
</html>
`

func simpleComparisonTest(t *testing.T, in, exp interface{}) {
	if in != exp {
		t.Errorf("Unexpected value, got %v", in)
	}
}

func TestExtractFormTokens(t *testing.T) {
	testScraper := new(scraper)
	err := testScraper.updateTokens(exampleLoginPage)
	if err != nil {
		t.Errorf("Got unexpected error: %v", err.Error())
	}
	t.Run("ViewState", func(t *testing.T) {
		correctViewState := "/wEPDwUKMTEzNTI2MjQyMg9kFgQCAw9kFgoCAQ9kFgICAQ8PFgIeBFRleHQFClBUTSBLb5J0bnJkZAICDxYCHgdWaXNpYmxlaGQCAkadl88BaGQCBA8WAh8BIdjwaQ8WAh8BaBYCAgEPEGRkFgBkAgUPDxYCHwAFCShkZXNrdG9wKWRkZMD6W4zuUf8/+Dp04d7fIVjURaNB"
		simpleComparisonTest(t, testScraper.viewState, correctViewState)
	})

	t.Run("ViewStateGenerator", func(t *testing.T) {
		correctViewStateGen := "CX0A0534"
		simpleComparisonTest(t, testScraper.viewStateGen, correctViewStateGen)
	})

	t.Run("EventValidation", func(t *testing.T) {
		correctEventValidation := "/wEWBAKNp6aXDwLi0uqnCgKF59rWBAK94jAbChDpRor7Q17FkzZGQUmFZy1hHJ2a"
		simpleComparisonTest(t, testScraper.eventValidation, correctEventValidation)
	})
}

func TestExtractCardData(t *testing.T) {
	data, err := parseData(exampleCardDetailsPage)
	if err != nil {
		t.Errorf("Got unexpected error: %v", err.Error())
	}
	t.Run("CardName", func(t *testing.T) {
		name := "John Doe"
		simpleComparisonTest(t, data.FullName, name)
	})

	t.Run("CardEmail", func(t *testing.T) {
		email := "john@example.com"
		simpleComparisonTest(t, data.Email, email)
	})

	t.Run("CardNumber", func(t *testing.T) {
		number := "2222333344445555"
		simpleComparisonTest(t, data.CardNumber, number)
	})

	t.Run("CardValue", func(t *testing.T) {
		balance := 69.42
		simpleComparisonTest(t, data.Balance, balance)
	})
}
