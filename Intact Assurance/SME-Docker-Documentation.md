# SME Docker Documentation

## Not supported on Internet Explorer
---

This application is used to easily deploy Guidewire environment. Each user has it's own independant environment that can be used for various purposes.

## How to use the app

1. To deploy your own Guidewire environment click onto this icon :

    <img src="deploy.png" alt="drawing" width="75"/>
2. Depending of your case choose the right tab and this will display the list of **pods** that are available.
    > A **pod** is an independant Guidewire environment that can be deploy on command. 

    Each item of the list is composed this way : <span style="text-decoration: underline;text-decoration-color: red;">SME-TRN-01</span> is the name of the pod, the <span style="text-decoration: underline;text-decoration-color:green;"> first URL</span> linked to the Guidewire environment and the <span style="text-decoration: underline;text-decoration-color: blue;">second</span> is to consult the reports associated with the specified environment : 
   
   <img src="item2.png" alt="drawing" width=""/>

   Click on the name of the pod, in this case <span style="color:red">SME-TRN-01</span> to show more details.
3. The first line in the pod menu tells some information about the pod and the second line allows to **deploy**, **restart** or **stop** a pod.
   <img src="menu.png" alt="drawing" width=""/>
  To deploy an environment: 
     1. check *Reset DB*
     2. choose a branch on the dropdown menu 
     3. click the *deploy* button

## Notes
>To restart the database of an environment, re-deploy it and make sure the *Reset DB* checkbox is checked.

>When you are done with an environment please **stop** it to relocate the ressources that are dedicated to it.

>The details concerning a pod do not update on their own. To see the status of a pod before or after performing an action always press the refresh button: <img src="refresh.png" alt="drawing" width="50"/> 

<br>
<br>

## Environments available in the app: 
> **SME environments**: Allow SME and BA to be able to test Guidewire Billing Center with any data and in any way they need without impacting any other person

> **SME-TR​N environments**: These environments are dedicated to Super Users so they can prepare for training sessions

> **Sandbox environments**: This environment contains the 'Out of the Box' version of Guidewire Billing Center for references

