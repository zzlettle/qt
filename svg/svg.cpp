#include "svg.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QGraphicsSvgItem>
#include <QIODevice>
#include <QMetaObject>
#include <QObject>
#include <QPaintEvent>
#include <QPainter>
#include <QRect>
#include <QRectF>
#include <QSize>
#include <QString>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionGraphicsItem>
#include <QSvgGenerator>
#include <QSvgRenderer>
#include <QSvgWidget>
#include <QWidget>
#include <QXmlStreamReader>

class MyQGraphicsSvgItem: public QGraphicsSvgItem {
public:
	void paint(QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget) { if (!callbackQGraphicsSvgItemPaint(this->objectName().toUtf8().data(), painter, const_cast<QStyleOptionGraphicsItem*>(option), widget)) { QGraphicsSvgItem::paint(painter, option, widget); }; };
protected:
};

char* QGraphicsSvgItem_ElementId(void* ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->elementId().toUtf8().data();
}

void QGraphicsSvgItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsSvgItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void* QGraphicsSvgItem_Renderer(void* ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->renderer();
}

void QGraphicsSvgItem_SetElementId(void* ptr, char* id){
	static_cast<QGraphicsSvgItem*>(ptr)->setElementId(QString(id));
}

void QGraphicsSvgItem_SetMaximumCacheSize(void* ptr, void* size){
	static_cast<QGraphicsSvgItem*>(ptr)->setMaximumCacheSize(*static_cast<QSize*>(size));
}

void QGraphicsSvgItem_SetSharedRenderer(void* ptr, void* renderer){
	static_cast<QGraphicsSvgItem*>(ptr)->setSharedRenderer(static_cast<QSvgRenderer*>(renderer));
}

int QGraphicsSvgItem_Type(void* ptr){
	return static_cast<QGraphicsSvgItem*>(ptr)->type();
}

class MyQSvgGenerator: public QSvgGenerator {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQSvgGenerator() : QSvgGenerator() {};
protected:
};

char* QSvgGenerator_Description(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->description().toUtf8().data();
}

char* QSvgGenerator_FileName(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->fileName().toUtf8().data();
}

void* QSvgGenerator_OutputDevice(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->outputDevice();
}

int QSvgGenerator_Resolution(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->resolution();
}

void QSvgGenerator_SetDescription(void* ptr, char* description){
	static_cast<QSvgGenerator*>(ptr)->setDescription(QString(description));
}

void QSvgGenerator_SetFileName(void* ptr, char* fileName){
	static_cast<QSvgGenerator*>(ptr)->setFileName(QString(fileName));
}

void QSvgGenerator_SetOutputDevice(void* ptr, void* outputDevice){
	static_cast<QSvgGenerator*>(ptr)->setOutputDevice(static_cast<QIODevice*>(outputDevice));
}

void QSvgGenerator_SetResolution(void* ptr, int dpi){
	static_cast<QSvgGenerator*>(ptr)->setResolution(dpi);
}

void QSvgGenerator_SetSize(void* ptr, void* size){
	static_cast<QSvgGenerator*>(ptr)->setSize(*static_cast<QSize*>(size));
}

void QSvgGenerator_SetTitle(void* ptr, char* title){
	static_cast<QSvgGenerator*>(ptr)->setTitle(QString(title));
}

void QSvgGenerator_SetViewBox(void* ptr, void* viewBox){
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRect*>(viewBox));
}

void QSvgGenerator_SetViewBox2(void* ptr, void* viewBox){
	static_cast<QSvgGenerator*>(ptr)->setViewBox(*static_cast<QRectF*>(viewBox));
}

char* QSvgGenerator_Title(void* ptr){
	return static_cast<QSvgGenerator*>(ptr)->title().toUtf8().data();
}

void* QSvgGenerator_NewQSvgGenerator(){
	return new MyQSvgGenerator();
}

void QSvgGenerator_DestroyQSvgGenerator(void* ptr){
	static_cast<QSvgGenerator*>(ptr)->~QSvgGenerator();
}

char* QSvgGenerator_ObjectNameAbs(void* ptr){
	return static_cast<MyQSvgGenerator*>(ptr)->objectNameAbs().toUtf8().data();
}

void QSvgGenerator_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQSvgGenerator*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQSvgRenderer: public QSvgRenderer {
public:
	void Signal_RepaintNeeded() { callbackQSvgRendererRepaintNeeded(this->objectName().toUtf8().data()); };
protected:
};

int QSvgRenderer_FramesPerSecond(void* ptr){
	return static_cast<QSvgRenderer*>(ptr)->framesPerSecond();
}

void QSvgRenderer_SetFramesPerSecond(void* ptr, int num){
	static_cast<QSvgRenderer*>(ptr)->setFramesPerSecond(num);
}

void QSvgRenderer_SetViewBox(void* ptr, void* viewbox){
	static_cast<QSvgRenderer*>(ptr)->setViewBox(*static_cast<QRect*>(viewbox));
}

void QSvgRenderer_SetViewBox2(void* ptr, void* viewbox){
	static_cast<QSvgRenderer*>(ptr)->setViewBox(*static_cast<QRectF*>(viewbox));
}

void* QSvgRenderer_NewQSvgRenderer(void* parent){
	return new QSvgRenderer(static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer4(void* contents, void* parent){
	return new QSvgRenderer(static_cast<QXmlStreamReader*>(contents), static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer3(void* contents, void* parent){
	return new QSvgRenderer(*static_cast<QByteArray*>(contents), static_cast<QObject*>(parent));
}

void* QSvgRenderer_NewQSvgRenderer2(char* filename, void* parent){
	return new QSvgRenderer(QString(filename), static_cast<QObject*>(parent));
}

int QSvgRenderer_Animated(void* ptr){
	return static_cast<QSvgRenderer*>(ptr)->animated();
}

int QSvgRenderer_ElementExists(void* ptr, char* id){
	return static_cast<QSvgRenderer*>(ptr)->elementExists(QString(id));
}

int QSvgRenderer_IsValid(void* ptr){
	return static_cast<QSvgRenderer*>(ptr)->isValid();
}

int QSvgRenderer_Load3(void* ptr, void* contents){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QXmlStreamReader*, static_cast<QXmlStreamReader*>(contents)));
}

int QSvgRenderer_Load2(void* ptr, void* contents){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QByteArray, *static_cast<QByteArray*>(contents)));
}

int QSvgRenderer_Load(void* ptr, char* filename){
	return QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "load", Q_ARG(QString, QString(filename)));
}

void QSvgRenderer_Render(void* ptr, void* painter){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)));
}

void QSvgRenderer_Render2(void* ptr, void* painter, void* bounds){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_Render3(void* ptr, void* painter, char* elementId, void* bounds){
	QMetaObject::invokeMethod(static_cast<QSvgRenderer*>(ptr), "render", Q_ARG(QPainter*, static_cast<QPainter*>(painter)), Q_ARG(QString, QString(elementId)), Q_ARG(QRectF, *static_cast<QRectF*>(bounds)));
}

void QSvgRenderer_ConnectRepaintNeeded(void* ptr){
	QObject::connect(static_cast<QSvgRenderer*>(ptr), static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), static_cast<MyQSvgRenderer*>(ptr), static_cast<void (MyQSvgRenderer::*)()>(&MyQSvgRenderer::Signal_RepaintNeeded));;
}

void QSvgRenderer_DisconnectRepaintNeeded(void* ptr){
	QObject::disconnect(static_cast<QSvgRenderer*>(ptr), static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), static_cast<MyQSvgRenderer*>(ptr), static_cast<void (MyQSvgRenderer::*)()>(&MyQSvgRenderer::Signal_RepaintNeeded));;
}

void QSvgRenderer_DestroyQSvgRenderer(void* ptr){
	static_cast<QSvgRenderer*>(ptr)->~QSvgRenderer();
}

class MyQSvgWidget: public QSvgWidget {
public:
	MyQSvgWidget(QWidget *parent) : QSvgWidget(parent) {};
	MyQSvgWidget(const QString &file, QWidget *parent) : QSvgWidget(file, parent) {};
protected:
	void paintEvent(QPaintEvent * event) { if (!callbackQSvgWidgetPaintEvent(this->objectName().toUtf8().data(), event)) { QSvgWidget::paintEvent(event); }; };
};

void* QSvgWidget_NewQSvgWidget(void* parent){
	return new MyQSvgWidget(static_cast<QWidget*>(parent));
}

void* QSvgWidget_NewQSvgWidget2(char* file, void* parent){
	return new MyQSvgWidget(QString(file), static_cast<QWidget*>(parent));
}

void QSvgWidget_Load2(void* ptr, void* contents){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QByteArray, *static_cast<QByteArray*>(contents)));
}

void QSvgWidget_Load(void* ptr, char* file){
	QMetaObject::invokeMethod(static_cast<QSvgWidget*>(ptr), "load", Q_ARG(QString, QString(file)));
}

void* QSvgWidget_Renderer(void* ptr){
	return static_cast<QSvgWidget*>(ptr)->renderer();
}

void QSvgWidget_DestroyQSvgWidget(void* ptr){
	static_cast<QSvgWidget*>(ptr)->~QSvgWidget();
}

