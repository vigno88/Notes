# Task anto

This is the method called in the screenDef with the name of the screenDef as argument

It returns an object of type BillingItaList

```java
public static BillingItaList createBillingListObj(String appId, String objListId)  throws ItaException {
		return new BillingItaList(appId, objListId, 0);
	}
```
This constructor call its parent class which is AbstractItaList_Table

```java
public BillingItaList(String appId, String objListId, int indexBased) throws ItaException {
	super(appId, objListId, indexBased);
```	
	
This constructor call one again its parent class which is AbstractItaList
```java
public AbstractItaList_Table(String screenId, String objListId, int indexBased) throws ItaException {
		super(screenId, objListId, indexBased);
	};
```	
This is the final constructor called. It sets the index (int), the screenDef (ScreenDefinition object), the listElement (ItaWebElement, fields object is a Map (key= field appId, value=ITA WebElement))

```java
public AbstractItaList(String screenId, String objListId, int indexBased) throws ItaException {
        // set index value which is 0 in our case
		this.indexBased = indexBased;
		
        // error handling
		if(StringUtils.isBlank(screenId)) {
			throw ItaMessageBundle.MESSAGES.illegalArgumentNullOrEmpty("screenId");
		}
		// get the screenDef object from the ItaRuntime
		this.screenDefinition = ItaRuntime.env.getScreenDefManager().getScreenDefinition(screenId);
		if(this.screenDefinition == null) {
			throw ItaMessageBundle.MESSAGES.screenNotFound(screenId);
		}

        // get the listElement ( ItaWebElement), this is the map of the fields in the webpage?
		this.listElement = findListWebElement(screenDefinition, objListId);
		if(this.listElement == null) {
			throw ItaMessageBundle.MESSAGES.webElementNotFoundInGivenScreen(!StringUtils.isBlank(objListId) ? objListId : "element with keyColumn true", screenId);
		}
		
		loadColumns();		
	}
```	

loadColumns set the number of columns in the screenDef and the name of all the field

This create the ItaList by invoking createMethod

```java
public static IItaList createNewList(String method) throws ItaException {
		if(StringUtils.isBlank(method)) {
			throw ItaMessageBundle.MESSAGES.illegalArgumentNullOrEmpty("method");
		}

		Matcher matcher = METHOD_PATTERN.matcher(method);		
		if(!matcher.find()) {
			throw ItaMessageBundle.MESSAGES.invalidListCreationMethodDeclaration(method);
		}
		
		String methodName = matcher.group(1);
		Object[] methodArguments = getMethodArgumentsArray(matcher.group(2));
		
		Method createMethod = getListClassByCreationMethod(methodName, methodArguments.length);
		try {
			return (IItaList)createMethod.invoke(null, methodArguments);
		} catch (SecurityException | IllegalAccessException | IllegalArgumentException | InvocationTargetException e) {
			throw ItaMessageBundle.MESSAGES.cannotCreateItaList(method, e);
		}
	}
```

```java
protected final IItaList createObjList(ScreenDefinition screenDef) throws ItaException {
		IItaList list = null;
		
		// if(sameAsPreviousAction(screenDef) && 
		// 		getPreviousAction().getRecordParameters().isExtraPropertyExist(OBJ_LIST_PARAM_NAME) &&
		// 		getPreviousAction().getRecordParameters().getExtraProperty(OBJ_LIST_PARAM_NAME, IItaList.class).getScreenDef().getAppId().equals(screenDef.getAppId())) {
			
		// 	list = getPreviousAction().getRecordParameters().getExtraProperty(OBJ_LIST_PARAM_NAME, IItaList.class);
		// } else {
			list = ItaListFactory.createNewList(screenDef);
			validateSelectedSheet(screenDef, list);
			list.refresh();
		}
		
		return list;
	}
```

This function is the one that will modify the screen, it gets the ItaList from the screenDef and retrieve the row that has to be modified from the action list

```java
@Keyword(value = Keywords.MODIFY_SCREEN, monitor = true, type = KeywordActionType.Modify)
	public boolean modify(ItaAction action) throws ItaException{
		action.selectDataSheet(true, true);
		
		ScreenDefinition screenDef = getScreenDefFromCurrentSheet(true);
		
		if(screenDef != null) {
			IItaList list = retrieveList(screenDef);
			ListItemRow listRow = retrieveRow(action, list);
			
			if(listRow != null){
				if(modifyItemInList(action, screenDef, list, listRow)) {
					return true;
				} else {
					if(ItaRuntime.env.getAppConfigurator().isTriggerFatalErrorOnModify()) {
						triggerFatalError();	
					}
					return false;
				}
			} else {
				return screenDef.populateValues(action);
			}
		}
		
		return false;
	}
	
```

This is where the field on the webpage is modified

``` java
private boolean modifyItemInList(ItaAction action, ScreenDefinition screenDef, IItaList list, ListItemRow listRow) throws ItaException {
		boolean modifyItemInList = true;
		
		getReportHandler().logInfoMessage("Update row " + (list.getSelectedRowIndex() +1));
		
		DataParameter[] parameters = getCurrentSheet().getParameters();
		
		int columnIndex;
		
		for (int i = 0; i < parameters.length; i++) {	
			if(!parameters[i].getValue().isNA() && !parameters[i].isCommentColumn()) {
				columnIndex = list.getColumnIndexByHeader(parameters[i].getName());
				
				if(columnIndex != -1) {
					try {
						if(!list.setCellValue(columnIndex, list.getSelectedRowIndex(), parameters[i].getValue().getEvaluatedValue(true))) {
							modifyItemInList = false;
						}
					} catch(Exception e) {
						getReportHandler().logErrorMessage("An error occurs when set value in field " + parameters[i].getName() + ".");
						getReportHandler().logErrorMessage(e);
						modifyItemInList = false;
					}
				}
			}
		}
		
		//Reset the selected row after the modify to force calling the row creation method
		list.resetSelectedRow();
		action.getRecordParameters().addExtraProperty(OBJ_LIST_PARAM_NAME, list);
		return modifyItemInList;
	}
```