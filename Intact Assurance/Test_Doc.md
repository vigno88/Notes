# Test Learning

List tasks :
- modify spring boot code
- modify web code
- using jenkins
- using maven
- using github
- using scrum
- 
## General Info

- [Jenkins Url](https://stha38e56:444/view/Billing/)

  - Once in the flow, we can see a list of build
    - BUILD_ENDTOEND_GWBC-BDI = gobrio (old system that we want to replace by Guidewire)
    - BUILD_ENDTOEND_GWBC = Guidewire
    - D1 <- D2 <- D3
- To see the result of the test int the ITA gui, click on **Results URL**
- Do not note in the suite follow up anything related to **Nouvelle Affaire**
- Doc [ITA](https://apps.iad.ca.inet/sites/adsth/ITA/default.aspx )
  
## Tasks to do in the morning

[3.Scheduled Tests [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/)

Always verify if the fail already happened in the past
BDI in priority since other test depends of them later in the day

> **BDI** username : look at the execution list 
>         password :dadev999

- [x] [BUILD_ENDTOEND_GWBC-GW_D1_FLOW](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-GW_D1_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-GW_D2_FLOW](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-GW_D2_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-GW_D3_FLOW](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-GW_D3_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-BDI_D1_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D1_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-BDI_D2_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D2_FLOW/)
- [ ] [BUILD_ENDTOEND_GWBC-BDI_D3_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_D3_FLOW/)
- [x] [BUILD_ENDTOEND_GWBC-BDI_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/BUILD_ENDTOEND_GWBC-BDI_FLOW/)
- [x] Look at the ping in the most recent build in [TEST_Integration [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_Integration/)
- [x] [TEST_SERVICES_JMeter [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TEST_SERVICES_JMeter/)

- [ ] [TESTS_ALL_GWBC_FLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_ALL_GWBC_FLOW/)
    > 1276, les rapport semble avoir rappetisser, le screendef a trop de column donc null pointer exception

    testReAssignDelinquencyBEL_BEL
- [ ] [TESTS_billing-BP-Clock [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-BP-Clock/)

- [x] [TESTS_billing-delinquency_part1 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part1/)
- [x] [TESTS_billing-delinquency_part2 [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_billing-delinquency_part2/)
  >  1621-DelinquencyCancellationRenewalNonResponsive_CL, pass
- [x] [TESTS_Billing-statement [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_Billing-statement/)
 > fail a ete relaunch
- [x] [TESTS_WORKFLOW [Jenkins]](https://stha38e56:444/view/Billing/view/3.Scheduled%20Tests/job/TESTS_WORKFLOW/)
---

## How to fill Testsuit follow-up

- Test Suite
  - Name of the suite in Jenkins
- Test failing
  - Can be found in the HTML report
- Branch
  - Can be found in the HTML report
- Date
  - suite dashboard
- Environment
  - **AppConfigEnv** in the dashboard of the test report
- Report
  - The url of the test report
- Execution host
  - **Hostname** in the test report
- Screen def
  - In the test report, second menu from the top, in the test scrollHonolulu folder as been deletedmenu, find the test that failed. In the description of the test, the screendef can be found in the step before the one that failed.
- Test fails since:
  - Date of the failure
- Detail:
  - Conclusion of the investigation

## Test avec proxy

603-DesktopActivities Didn't fail
1008-SearchActivities Didn't fail

To check next time :

 356-ActivityEnhAcc 

## Task

LOAD_BILLING_STATEMENT_OUTPUT_FILE
SELECT_BILLING_STATEMENT_POLICY_PERIOD
VALIDATE_BILLING_STATEMENT_RECORD


<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:soap="http://guidewire.com/ws/soapheaders" xmlns:bil="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BillingDocumentAPI">  
<soapenv:Header>
      <soap:authentication>
         <soap:username>su</soap:username>
         <soap:password>gw</soap:password>
      </soap:authentication>
   </soapenv:Header>
   <soapenv:Body>
      <bil:generateBillingStatementInfo>
         <!--Optional:-->
         <bil:billingStatementReferenceNumber>[parameter1]</bil:billingStatementReferenceNumber>
         <!--Optional:-->
         <bil:contractReferenceNumber></bil:contractReferenceNumber>
         <!--Optional:-->
         <bil:requestor></bil:requestor>
         <!--Optional:-->
         <bil:policyNumber>[parameter2]</bil:policyNumber>
      </bil:generateBillingStatementInfo>
   </soapenv:Body>
</soapenv:Envelope>

<?xml version="1.0"?>
<tns:Envelope xmlns:tns="http://schemas.xmlsoap.org/soap/envelope/">
  <tns:Body>
    <generateBillingStatementInfoResponse xmlns="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BillingDocumentAPI">
      <return xmlns:pogo="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCBillingStatementInfo_IFC">
        <pogo:AccountNumber>1000000016</pogo:AccountNumber>
        <pogo:AccountPreviousTotalBalance>0</pogo:AccountPreviousTotalBalance>
        <DefaultPaymentInstrument xmlns="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCBillingStatementInfo_IFC" xmlns:pogo="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCPaymentInstrument_IFC">
          <pogo:CreditCardAccountNumber>nullXXXXXXXXnull</pogo:CreditCardAccountNumber>
          <pogo:MaskedFinancialInstAccountNumber>XXXX927</pogo:MaskedFinancialInstAccountNumber>
          <pogo:PaymentMethod>ach</pogo:PaymentMethod>
        </DefaultPaymentInstrument>
        <pogo:InsufficientFundsFees>0</pogo:InsufficientFundsFees>
        <pogo:PaymentDueAmount>171.74</pogo:PaymentDueAmount>
        <pogo:PaymentDueDate>2019-07-24T00:00:00-04:00</pogo:PaymentDueDate>
        <pogo:PaymentPlan>mo</pogo:PaymentPlan>
        <pogo:PolicyReinstatefees>0.00</pogo:PolicyReinstatefees>
        <pogo:PolicyUnapplied>0.00</pogo:PolicyUnapplied>
        <pogo:RemainingBalance>2060.00</pogo:RemainingBalance>
        <pogo:RemainingInstallmentsCount>12</pogo:RemainingInstallmentsCount>
        <pogo:StatementInsufficientFundsCharges>0</pogo:StatementInsufficientFundsCharges>
        <pogo:StatementTaxAmount>0</pogo:StatementTaxAmount>
        <pogo:TotalBankFees>0</pogo:TotalBankFees>
        <pogo:TotalFinancialFees>60.00</pogo:TotalFinancialFees>
        <pogo:TotalPostDatedChequeAmount>0</pogo:TotalPostDatedChequeAmount>
        <pogo:TransactionsList>
          <Entry xmlns="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCBillingStatementInfo_IFC" xmlns:pogo="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCBillingStatementTransaction_IFC">
            <pogo:Code>NBS</pogo:Code>
            <pogo:EffectiveDate>2019-07-22T00:00:00-04:00</pogo:EffectiveDate>
            <pogo:FeesAmount>60.00</pogo:FeesAmount>
            <pogo:PremiumAmount>2000.00</pogo:PremiumAmount>
            <pogo:PrincipalTransactionIndicator>Y</pogo:PrincipalTransactionIndicator>
            <pogo:TotalAmount>2060.00</pogo:TotalAmount>
          </Entry>
          <Entry xmlns="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCBillingStatementInfo_IFC" xmlns:pogo="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCBillingStatementTransaction_IFC">
            <pogo:Code>NBS</pogo:Code>
            <pogo:EffectiveDate>2019-07-22T00:00:00-04:00</pogo:EffectiveDate>
            <pogo:FeesAmount>60.00</pogo:FeesAmount>
            <pogo:PremiumAmount>2000.00</pogo:PremiumAmount>
            <pogo:PrincipalTransactionIndicator>Y</pogo:PrincipalTransactionIndicator>
            <pogo:TotalAmount>2060.00</pogo:TotalAmount>
          </Entry>
          <Entry xmlns="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCBillingStatementInfo_IFC" xmlns:pogo="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCBillingStatementTransaction_IFC">
            <pogo:Code>NBS</pogo:Code>
            <pogo:EffectiveDate>2019-07-22T00:00:00-04:00</pogo:EffectiveDate>
            <pogo:FeesAmount>60.00</pogo:FeesAmount>
            <pogo:PremiumAmount>2000.00</pogo:PremiumAmount>
            <pogo:PrincipalTransactionIndicator>Y</pogo:PrincipalTransactionIndicator>
            <pogo:TotalAmount>2060.00</pogo:TotalAmount>
          </Entry>
          <Entry xmlns="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCBillingStatementInfo_IFC" xmlns:pogo="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BCBillingStatementTransaction_IFC">
            <pogo:Code>NBS</pogo:Code>
            <pogo:EffectiveDate>2019-07-22T00:00:00-04:00</pogo:EffectiveDate>
            <pogo:FeesAmount>60.00</pogo:FeesAmount>
            <pogo:PremiumAmount>2000.00</pogo:PremiumAmount>
            <pogo:PrincipalTransactionIndicator>Y</pogo:PrincipalTransactionIndicator>
            <pogo:TotalAmount>2060.00</pogo:TotalAmount>
          </Entry>

        </pogo:TransactionsList>
      </return>
    </generateBillingStatementInfoResponse>
  </tns:Body>
</tns:Envelope>





<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:soap="http://guidewire.com/ws/soapheaders" xmlns:bil="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BillingDocumentAPI">  
<soapenv:Header>
      <soap:authentication>
         <soap:username>su</soap:username>
         <soap:password>gw</soap:password>
      </soap:authentication>
   </soapenv:Header>
   <soapenv:Body>
      <bil:generateBillingStatementInfo>
         <!--Optional:-->
         <bil:billingStatementReferenceNumber>[parameter1]</bil:billingStatementReferenceNumber>
         <!--Optional:-->
         <bil:contractReferenceNumber></bil:contractReferenceNumber>
         <!--Optional:-->
         <bil:requestor></bil:requestor>
         <!--Optional:-->
         <bil:policyNumber>[parameter2]</bil:policyNumber>
      </bil:generateBillingStatementInfo>
   </soapenv:Body>
</soapenv:Envelope>

<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope" xmlns:soap="http://guidewire.com/ws/soapheaders" xmlns:bil="http://guidewire.com/bc/ws/ifc/webservice/document/bc801/BillingDocumentAPI">
  <soapenv:Header>
    <soap:authentication>
      <soap:username>su</soap:username>
      <soap:password>gw</soap:password>
    </soap:authentication>
  </soapenv:Header>
  <soapenv:Body>
    <bil:generateBillingStatementInfo>
      <!--Optional:-->
      <bil:billingStatementReferenceNumber>GWBC2</bil:billingStatementReferenceNumber><!--Optional:--><bil:contractReferenceNumber></bil:contractReferenceNumber><!--Optional:--><bil:requestor></bil:requestor><!--Optional:--><bil:policyNumber>B00-0003</bil:policyNumber></bil:generateBillingStatementInfo></soapenv:Body></soapenv:Envelope>