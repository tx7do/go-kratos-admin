/**
 * User: Jinqn
 * Date: 14-04-08
 * Time: 下午16:34
 * 上传图片对话框逻辑代码,包括tab: 远程图片/上传图片/在线图片/搜索图片
 */
(function () {

  let onlineImage,
    remoteImage,
    uploadImage;
  let editorOpt = {};

  window.addEventListener('load', () => {
    editorOpt = editor.getOpt('imageConfig');
    initTabs();
    initAlign();
    initButtons();
  });

  /* 初始化tab标签 */
  function initTabs() {
    const tabs = $G('tabhead').children;
    for (const tab of tabs) {
      domUtils.on(tab, "click", (e) => {
        const target = e.target || e.srcElement;
        setTabFocus(target.dataset.contentId);
      });
    }
    if (!editorOpt.disableUpload) {
      $G('tabhead').querySelector('[data-content-id="upload"]').style.display = 'inline-block';
    }
    if (!editorOpt.disableOnline) {
      $G('tabhead').querySelector('[data-content-id="online"]').style.display = 'inline-block';
    }
    if (editorOpt.selectCallback) {
      $G('imageSelect').style.display = 'inline-block';
      domUtils.on($G('imageSelect'), "click", (e) => {
        editorOpt.selectCallback(editor, (info) => {
          if (info) {
            $G('url').value = info.path;
            $G('title').value = info.name;
            const img = new Image();
            img.addEventListener('load', () => {
              $G('width').value = img.width;
              $G('height').value = img.height;
              remoteImage.setPreview();
            });
            img.onerror = function () {
              remoteImage.setPreview();
            };
            img.src = info.path;
          }
        });
      });
    }
    const img = editor.selection.getRange().getClosedNode();
    if (img && img.tagName && img.tagName.toLowerCase() == 'img') {
      setTabFocus('remote');
    } else {
      setTabFocus('remote');
    }
  }

  /* 初始化tabbody */
  function setTabFocus(id) {
    if (!id) return;
    let bodyId; let i; const tabs = $G('tabhead').children;
    for (i = 0; i < tabs.length; i++) {
      bodyId = tabs[i].dataset.contentId;
      if (bodyId == id) {
        domUtils.addClass(tabs[i], 'focus');
        domUtils.addClass($G(bodyId), 'focus');
      } else {
        domUtils.removeClasses(tabs[i], 'focus');
        domUtils.removeClasses($G(bodyId), 'focus');
      }
    }
    switch (id) {
      case 'online': {
        setAlign(editor.getOpt('imageManagerInsertAlign'));
        onlineImage = onlineImage || new OnlineImage('imageList');
        onlineImage.reset();
        break;
      }
      case 'remote': {
        remoteImage = remoteImage || new RemoteImage();
        break;
      }
      case 'upload': {
        setAlign(editor.getOpt('imageInsertAlign'));
        uploadImage = uploadImage || new UploadImage('queueList');
        break;
      }
    }
  }

  /* 初始化onok事件 */
  function initButtons() {

    dialog.onok = function () {
      let id; let list = []; const remote = false; const tabs = $G('tabhead').children;
      for (const tab of tabs) {
        if (domUtils.hasClass(tab, 'focus')) {
          id = tab.dataset.contentId;
          break;
        }
      }

      switch (id) {
        case 'online': {
          list = onlineImage.getInsertList();
          break;
        }
        case 'remote': {
          list = remoteImage.getInsertList();
          break;
        }
        case 'upload': {
          list = uploadImage.getInsertList();
          const count = uploadImage.getQueueCount();
          if (count) {
            $('.info', '#queueList').html(`<span style="color:red;">${  '还有2个未上传文件'.replace(/\d/, count)  }</span>`);
            return false;
          }
          break;
        }
      }

      if (list) {
        editor.execCommand('insertimage', list);
        remote && editor.fireEvent("catchRemoteImage");
      }
    };
  }


  /* 初始化对其方式的点击事件 */
  function initAlign() {
    /* 点击align图标 */
    domUtils.on($G("alignIcon"), 'click', (e) => {
      const target = e.target || e.srcElement;
      if (target.className && target.className.includes('-align')) {
        setAlign(target.dataset.align);
      }
    });
  }

  /* 设置对齐方式 */
  function setAlign(align) {
    align = align || 'none';
    const aligns = $G("alignIcon").children;
    for (i = 0; i < aligns.length; i++) {
      if (aligns[i].dataset.align == align) {
        domUtils.addClass(aligns[i], 'focus');
        $G("align").value = aligns[i].dataset.align;
      } else {
        domUtils.removeClasses(aligns[i], 'focus');
      }
    }
  }

  /* 获取对齐方式 */
  function getAlign() {
    const align = $G("align").value || 'none';
    return align == 'none' ? '' : align;
  }


  /* 在线图片 */
  function RemoteImage(target) {
    this.container = utils.isString(target) ? document.getElementById(target) : target;
    this.init();
  }

  RemoteImage.prototype = {
    init () {
      this.initContainer();
      this.initEvents();
    },
    initContainer () {
      this.dom = {
        'url': $G('url'),
        'width': $G('width'),
        'height': $G('height'),
        'border': $G('border'),
        'vhSpace': $G('vhSpace'),
        'title': $G('title'),
        'align': $G('align')
      };
      const img = editor.selection.getRange().getClosedNode();
      if (img) {
        this.setImage(img);
      }
    },
    initEvents () {
      const _this = this;
        const locker = $G('lock');

      /* 改变url */
      domUtils.on($G("url"), 'keyup', updatePreview);
      domUtils.on($G("border"), 'keyup', updatePreview);
      domUtils.on($G("title"), 'keyup', updatePreview);

      domUtils.on($G("width"), 'keyup', function () {
        if (locker.checked) {
          const proportion = locker.dataset.proportion;
          $G('height').value = Math.round(this.value / proportion);
        } else {
          _this.updateLocker();
        }
        updatePreview();
      });
      domUtils.on($G("height"), 'keyup', function () {
        if (locker.checked) {
          const proportion = locker.dataset.proportion;
          $G('width').value = Math.round(this.value * proportion);
        } else {
          _this.updateLocker();
        }
        updatePreview();
      });
      domUtils.on($G("lock"), 'change', () => {
        const proportion = Number.parseInt($G("width").value) / Number.parseInt($G("height").value);
        locker.dataset.proportion = proportion;
      });

      function updatePreview() {
        _this.setPreview();
      }
    },
    updateLocker () {
      const height = $G('height').value;
        const locker = $G('lock');
        const width = $G('width').value;
      if (width && height && width == Number.parseInt(width) && height == Number.parseInt(height)) {
        locker.disabled = false;
        locker.title = '';
      } else {
        locker.checked = false;
        locker.disabled = 'disabled';
        locker.title = lang.remoteLockError;
      }
    },
    setImage (img) {
      /* 不是正常的图片 */
      if (!img.tagName || img.tagName.toLowerCase() != 'img' && !img.getAttribute("src") || !img.src) return;

      const align = editor.queryCommandValue("imageFloat");
        const wordImgFlag = img.dataset.wordImage;
        const src = wordImgFlag ? wordImgFlag.replace("&amp;", "&") : (img.getAttribute('_src') || img.getAttribute("src", 2).replace("&amp;", "&"));

      /* 防止onchange事件循环调用 */
      if (src !== $G("url").value) $G("url").value = src;
      if (src) {
        /* 设置表单内容 */
        $G("width").value = img.width || '';
        $G("height").value = img.height || '';
        $G("border").value = img.getAttribute("border") || '0';
        $G("vhSpace").value = img.getAttribute("vspace") || '0';
        $G("title").value = img.title || img.alt || '';
        setAlign(align);
        this.setPreview();
        this.updateLocker();
      }
    },
    getData () {
      const data = {};
      for (const k in this.dom) {
        data[k] = this.dom[k].value;
      }
      return data;
    },
    setPreview () {
      const border = $G('border').value;
        let height;
        const oh = $G('height').value;
        const ow = $G('width').value;
        const preview = $G('preview');
        const title = $G('title').value;
        const url = $G('url').value;
        let width;

      width = ((!ow || !oh) ? preview.offsetWidth : Math.min(ow, preview.offsetWidth));
      width = width + (border * 2) > preview.offsetWidth ? width : (preview.offsetWidth - (border * 2));
      height = (!ow || !oh) ? '' : width * oh / ow;

      if (url) {
        preview.innerHTML = `<img src="${  url  }" width="${  width  }" height="${  height  }" border="${  border  }px solid #000" title="${  title  }" />`;
      }
    },
    getInsertList () {
      const data = this.getData();
      if (data.url) {
        const img = {
          src: data.url,
          _src: data.url,
        }
        img._propertyDelete = []
        img.style = []
        if (data.width) {
          img.width = data.width;
          img.style.push(`width:${  data.width  }px`);
        } else {
          img._propertyDelete.push('width');
        }
        if (data.height) {
          img.height = data.height;
          img.style.push(`height:${  data.height  }px`);
        } else {
          img._propertyDelete.push('height');
        }
        if (data.border) {
          img.border = data.border;
        } else {
          img._propertyDelete.push('border');
        }
        if (data.align) {
          img.floatStyle = data.align;
        } else {
          img._propertyDelete.push('floatStyle');
        }
        if (data.vhSpace) {
          img.vspace = data.vhSpace;
        } else {
          img._propertyDelete.push('vspace');
        }
        if (data.title) {
          img.alt = data.title;
        } else {
          img._propertyDelete.push('alt');
        }
        if (img.style.length > 0) {
          img.style = img.style.join(';');
        } else {
          img._propertyDelete.push('style');
        }
        return [img];
      } else {
        return [];
      }
    }
  };


  /* 上传图片 */
  function UploadImage(target) {
    this.$wrap = target.constructor == String ? $(`#${  target}`) : $(target);
    this.init();
  }

  UploadImage.prototype = {
    init () {
      this.imageList = [];
      this.initContainer();
      this.initUploader();
    },
    initContainer () {
      this.$queue = this.$wrap.find('.filelist');
    },
    /* 初始化容器 */
    initUploader () {
      const $ = jQuery;    // just in case. Make sure it's not an other libaray.
        const _this = this;
        const $wrap = _this.$wrap;
        // 上传按钮
        const $filePickerBlock = $wrap.find('.filePickerBlock');
        // 上传按钮
        const $filePickerBtn = $wrap.find('.filePickerBtn');
        // 状态栏，包括进度和控制按钮
        const $statusBar = $wrap.find('.statusBar');
        // 文件总体选择信息。
        const $info = $statusBar.find('.info');
        // 没选择文件之前的内容。
        const $placeHolder = $wrap.find('.placeholder');
        // 总体进度条
        const $progress = $statusBar.find('.progress').hide();
        // 图片容器
        const $queue = $wrap.find('.filelist');
        // 上传按钮
        const $upload = $wrap.find('.uploadBtn');
        const acceptExtensions = (editor.getOpt('imageAllowFiles') || []).join('').replaceAll('.', ',').replace(/^,/, '');
        const actionUrl = editor.getActionUrl(editor.getOpt('imageActionName'));
        // 添加的文件数量
        let fileCount = 0;
        // 添加的文件总大小
        let fileSize = 0;
        const imageCompressBorder = editor.getOpt('imageCompressBorder');
        const imageMaxSize = editor.getOpt('imageMaxSize');
        // 所有文件的进度信息，key为file id
        const percentages = {};
        // 优化retina, 在retina下这个值是2
        const ratio = window.devicePixelRatio || 1;
        // 可能有pedding, ready, uploading, confirm, done.
        let state = '';
        const supportTransition = (function () {
          let s = document.createElement('p').style;
            const r = 'transition' in s ||
              'WebkitTransition' in s ||
              'MozTransition' in s ||
              'msTransition' in s ||
              'OTransition' in s;
          s = null;
          return r;
        })();
        const thumbnailHeight = 113 * ratio;
        // 缩略图大小
        const thumbnailWidth = 113 * ratio;
        // WebUploader实例
        let uploader;

      if (!WebUploader.Uploader.support()) {
        $('#filePickerReady').after($('<div>').html(lang.errorNotSupport)).hide();
        return;
      } else if (!editor.getOpt('imageActionName')) {
        $('#filePickerReady').after($('<div>').html(lang.errorLoadConfig)).hide();
        return;
      }

      const uploaderOption = {
        pick: {
          id: '#filePickerReady',
          label: lang.uploadSelectFile
        },
        accept: {
          title: 'Images',
          extensions: acceptExtensions,
          mimeTypes: 'image/*'
        },
        swf: '../../third-party/webuploader/Uploader.swf',
        server: actionUrl,
        fileVal: editor.getOpt('imageFieldName'),
        duplicate: true,
        fileSingleSizeLimit: imageMaxSize,    // 默认 2 M
        threads: 1,
        headers: editor.getOpt('serverHeaders') || {},
        compress: editor.getOpt('imageCompressEnable') ? {
          enable: editor.getOpt('imageCompressEnable'),
          maxWidthOrHeight: imageCompressBorder,
          maxSize: imageMaxSize,
        } : false
      };
      if(editor.getOpt('uploadServiceEnable')) {
        uploaderOption.customUpload = function (file, callback) {
          editor.getOpt('uploadServiceUpload')('image', file, {
            success( res ) {
              callback.onSuccess(file, {_raw:JSON.stringify(res)});
            },
            error( err ) {
              callback.onError(file, err);
            },
            progress( percent ) {
              callback.onProgress(file, percent);
            }
          }, {
            from: 'image'
          });
        };
      }

      uploader = _this.uploader = WebUploader.create(uploaderOption);
      uploader.addButton({
        id: '#filePickerBlock'
      });
      uploader.addButton({
        id: '#filePickerBtn',
        label: lang.uploadAddFile
      });

      setState('pedding');

      // 当有文件添加进来时执行，负责view的创建
      function addFile(file) {
        const $li = $(`<li id="${  file.id  }">` +
            `<p class="title">${  file.name  }</p>` +
            `<p class="imgWrap"></p>` +
            `<p class="progress"><span></span></p>` +
            `</li>`);

          const $btns = $(`<div class="file-panel">` +
            `<span class="cancel">${  lang.uploadDelete  }</span>` +
            `<span class="rotateRight">${  lang.uploadTurnRight  }</span>` +
            `<span class="rotateLeft">${  lang.uploadTurnLeft  }</span></div>`).appendTo($li);
          const $info = $('<p class="error"></p>').hide().appendTo($li);
          const $prgress = $li.find('p.progress span');
          const $wrap = $li.find('p.imgWrap');

          const showError = function (code) {
            switch (code) {
              case 'exceed_size': {
                text = lang.errorExceedSize;
                break;
              }
              case 'http': {
                text = lang.errorHttp;
                break;
              }
              case 'interrupt': {
                text = lang.errorInterrupt;
                break;
              }
              case 'not_allow_type': {
                text = lang.errorFileType;
                break;
              }
              default: {
                text = lang.errorUploadRetry;
                break;
              }
            }
            $info.text(text).show();
          };

        if (file.getStatus() === 'invalid') {
          showError(file.statusText);
        } else {
          $wrap.text(lang.uploadPreview);
          if (browser.ie && browser.version <= 7) {
            $wrap.text(lang.uploadNoPreview);
          } else {
            uploader.makeThumb(file, (error, src) => {
              if (error || !src) {
                $wrap.text(lang.uploadNoPreview);
              } else {
                const $img = $(`<img src="${  src  }">`);
                $wrap.empty().append($img);
                $img.on('error', () => {
                  $wrap.text(lang.uploadNoPreview);
                });
              }
            }, thumbnailWidth, thumbnailHeight);
          }
          percentages[file.id] = [file.size, 0];
          file.rotation = 0;

          /* 检查文件格式 */
          if (!file.ext || !acceptExtensions.includes(file.ext.toLowerCase())) {
            showError('not_allow_type');
            uploader.removeFile(file);
          }
        }

        file.on('statuschange', (cur, prev) => {
          if (prev === 'progress') {
            $prgress.hide().width(0);
          } else if (prev === 'queued') {
            $li.off('mouseenter mouseleave');
            $btns.remove();
          }
          // 成功
          switch (cur) {
          case 'complete': {

          break;
          }
          case 'error':
          case 'invalid': {
            showError(file.statusText);
            percentages[file.id][1] = 1;

          break;
          }
          case 'interrupt': {
            showError('interrupt');

          break;
          }
          case 'progress': {
            $info.hide();
            $prgress.css('display', 'block');

          break;
          }
          case 'queued': {
            percentages[file.id][1] = 0;

          break;
          }
          // No default
          }

          $li.removeClass(`state-${  prev}`).addClass(`state-${  cur}`);
        });

        $li.on('mouseenter', () => {
          $btns.stop().animate({height: 30});
        });
        $li.on('mouseleave', () => {
          $btns.stop().animate({height: 0});
        });

        $btns.on('click', 'span', function () {
          let deg;
            const index = $(this).index();

          switch (index) {
            case 0: {
              uploader.removeFile(file);
              return;
            }
            case 1: {
              file.rotation += 90;
              break;
            }
            case 2: {
              file.rotation -= 90;
              break;
            }
          }

          if (supportTransition) {
            deg = `rotate(${  file.rotation  }deg)`;
            $wrap.css({
              '-webkit-transform': deg,
              '-mos-transform': deg,
              '-o-transform': deg,
              'transform': deg
            });
          } else {
            $wrap.css('filter', `progid:DXImageTransform.Microsoft.BasicImage(rotation=${  Math.trunc((file.rotation / 90) % 4 + 4) % 4  })`);
          }

        });

        $li.insertBefore($filePickerBlock);
      }

      // 负责view的销毁
      function removeFile(file) {
        const $li = $(`#${  file.id}`);
        delete percentages[file.id];
        updateTotalProgress();
        $li.off().find('.file-panel').off().end().remove();
      }

      function updateTotalProgress() {
        let loaded = 0;
          let percent;
          const spans = $progress.children();
          let total = 0;

        $.each(percentages, (k, v) => {
          total += v[0];
          loaded += v[0] * v[1];
        });

        percent = total ? loaded / total : 0;

        spans.eq(0).text(`${Math.round(percent * 100)  }%`);
        spans.eq(1).css('width', `${Math.round(percent * 100)  }%`);
        updateStatus();
      }

      function setState(val, files) {

        if (val !== state) {

          let stats = uploader.getStats();

          $upload.removeClass(`state-${  state}`);
          $upload.addClass(`state-${  val}`);

          switch (val) {

            case 'confirm': {
              $progress.show();
              $info.hide();
              $upload.text(lang.uploadStart);

              stats = uploader.getStats();
              if (stats.successNum && !stats.uploadFailNum) {
                setState('finish');
                return;
              }
              break;
            }

            case 'finish': {
              $progress.hide();
              $info.show();
              if (stats.uploadFailNum) {
                $upload.text(lang.uploadRetry);
              } else {
                $upload.text(lang.uploadStart);
              }
              break;
            }

            /* 暂停上传 */
            case 'paused': {
              $progress.show();
              $info.hide();
              $upload.text(lang.uploadContinue);
              break;
            }

            /* 未选择文件 */
            case 'pedding': {
              $queue.addClass('element-invisible');
              $statusBar.addClass('element-invisible');
              $placeHolder.removeClass('element-invisible');
              $progress.hide();
              $info.hide();
              uploader.refresh();
              break;
            }

            /* 可以开始上传 */
            case 'ready': {
              $placeHolder.addClass('element-invisible');
              $queue.removeClass('element-invisible');
              $statusBar.removeClass('element-invisible');
              $progress.hide();
              $info.show();
              $upload.text(lang.uploadStart);
              uploader.refresh();
              break;
            }

            /* 上传中 */
            case 'uploading': {
              $progress.show();
              $info.hide();
              $upload.text(lang.uploadPause);
              break;
            }
          }

          state = val;
          updateStatus();

        }

        if (_this.getQueueCount()) {
          $upload.removeClass('disabled')
        } else {
          $upload.addClass('disabled')
        }

      }

      function updateStatus() {
        let stats; let text = '';

        if (state === 'ready') {
          text = lang.updateStatusReady.replace('_', fileCount).replace('_KB', WebUploader.formatSize(fileSize));
        } else if (state === 'confirm') {
          stats = uploader.getStats();
          if (stats.uploadFailNum) {
            text = lang.updateStatusConfirm.replace('_', stats.successNum).replace('_', stats.successNum);
          }
        } else {
          stats = uploader.getStats();
          text = lang.updateStatusFinish.replace('_', fileCount).replace('_KB', WebUploader.formatSize(fileSize)).replace('_', stats.successNum);

          if (stats.uploadFailNum) {
            text += lang.updateStatusError.replace('_', stats.uploadFailNum);
          }
        }

        $info.html(text);
      }

      uploader.on('fileQueued', (file) => {
        fileCount++;
        fileSize += file.size;

        if (fileCount === 1) {
          $placeHolder.addClass('element-invisible');
          $statusBar.show();
        }

        addFile(file);
      });

      uploader.on('fileDequeued', (file) => {
        if (file.ext && acceptExtensions.includes(file.ext.toLowerCase()) && file.size <= imageMaxSize) {
          fileCount--;
          fileSize -= file.size;
        }

        removeFile(file);
        updateTotalProgress();
      });

      uploader.on('filesQueued', (file) => {
        if (!uploader.isInProgress() && (state == 'pedding' || state == 'finish' || state == 'confirm' || state == 'ready')) {
          setState('ready');
        }
        updateTotalProgress();
      });

      uploader.on('all', (type, files) => {
        switch (type) {
          case 'startUpload': {
            /* 添加额外的GET参数 */
            const params = utils.serializeParam(editor.queryCommandValue('serverparam')) || '';
              const url = utils.formatUrl(`${actionUrl + (actionUrl.includes('?') ? '&' : '?')  }encode=utf-8&${  params}`);
            uploader.option('server', url);
            setState('uploading', files);
            break;
          }
          case 'stopUpload': {
            setState('paused', files);
            break;
          }
          case 'uploadFinished': {
            setState('confirm', files);
            break;
          }
        }
      });

      uploader.on('uploadBeforeSend', (file, data, header) => {
        // 这里可以通过data对象添加POST参数
        if (actionUrl.toLowerCase().includes('jsp')) {
          header['X-Requested-With'] = 'XMLHttpRequest';
        }
      });

      uploader.on('uploadProgress', (file, percentage) => {
        const $li = $(`#${  file.id}`);
          const $percent = $li.find('.progress span');

        $percent.css('width', `${percentage * 100  }%`);
        percentages[file.id][1] = percentage;
        updateTotalProgress();
      });

      uploader.on('uploadSuccess', (file, ret) => {
        const $file = $(`#${  file.id}`);
        try {
          const responseText = (ret._raw || ret);
            let json = utils.str2json(responseText);
          json = editor.options.serverResponsePrepare(json);
          if (json.state == 'SUCCESS') {
            _this.imageList.push(json);
            $file.append('<span class="success"></span>');
            // 触发上传图片事件
            editor.fireEvent("uploadsuccess", {
              res: json,
              type: 'image'
            });
          } else {
            $file.find('.error').text(json.state).show();
          }
        } catch {
          $file.find('.error').text(lang.errorServerUpload).show();
        }
      });

      uploader.on('uploadError', (file, code) => {});
      uploader.on('error', (code, param1, param2) => {
        if (code === 'F_EXCEED_SIZE') {
          editor.getOpt('tipError')(`${lang.errorExceedSize  } ${  (param1 / 1024 / 1024).toFixed(1)  }MB`);
        } else {
          console.log('error', code, param1, param2);
        }
      });
      uploader.on('uploadComplete', (file, ret) => {});

      $upload.on('click', function () {
        if ($(this).hasClass('disabled')) {
          return false;
        }

        switch (state) {
        case 'paused': {
          uploader.upload();

        break;
        }
        case 'ready': {
          uploader.upload();

        break;
        }
        case 'uploading': {
          uploader.stop();

        break;
        }
        // No default
        }
      });

      $upload.addClass(`state-${  state}`);
      updateTotalProgress();
    },
    getQueueCount () {
      let file; const files = this.uploader.getFiles(); let i; let readyFile = 0; let status;
      for (i = 0; file = files[i++];) {
        status = file.getStatus();
        if (status == 'queued' || status == 'uploading' || status == 'progress') readyFile++;
      }
      return readyFile;
    },
    destroy () {
      this.$wrap.remove();
    },
    getInsertList () {
      const align = getAlign(); let data; let i;
        const list = [];
        const prefix = editor.getOpt('imageUrlPrefix');
      for (i = 0; i < this.imageList.length; i++) {
        data = this.imageList[i];
        list.push({
          src: prefix + data.url,
          _src: prefix + data.url,
          alt: data.original,
          floatStyle: align
        });
      }
      return list;
    }
  };


  /* 在线图片 */
  function OnlineImage(target) {
    this.container = utils.isString(target) ? document.getElementById(target) : target;
    this.init();
  }

  OnlineImage.prototype = {
    init () {
      this.reset();
      this.initEvents();
    },
    /* 初始化容器 */
    initContainer () {
      this.container.innerHTML = '';
      this.list = document.createElement('ul');
      this.clearFloat = document.createElement('li');

      domUtils.addClass(this.list, 'list');
      domUtils.addClass(this.clearFloat, 'clearFloat');

      this.list.append(this.clearFloat);
      this.container.append(this.list);
    },
    /* 初始化滚动事件,滚动到地步自动拉取数据 */
    initEvents () {
      const _this = this;

      /* 滚动拉取图片 */
      domUtils.on($G('imageList'), 'scroll', function (e) {
        const panel = this;
        if (panel.scrollHeight - (panel.offsetHeight + panel.scrollTop) < 10) {
          _this.getImageData();
        }
      });
      /* 选中图片 */
      domUtils.on(this.container, 'click', (e) => {
        const target = e.target || e.srcElement;
          const li = target.parentNode;

        if (li.tagName.toLowerCase() == 'li') {
          if (domUtils.hasClass(li, 'selected')) {
            domUtils.removeClasses(li, 'selected');
          } else {
            domUtils.addClass(li, 'selected');
          }
        }
      });
    },
    /* 初始化第一次的数据 */
    initData () {

      /* 拉取数据需要使用的值 */
      this.state = 0;
      this.listSize = editor.getOpt('imageManagerListSize');
      this.listIndex = 0;
      this.listEnd = false;

      /* 第一次拉取数据 */
      this.getImageData();
    },
    /* 重置界面 */
    reset () {
      this.initContainer();
      this.initData();
    },
    /* 向后台拉取图片列表数据 */
    getImageData () {
      const _this = this;

      if (!_this.listEnd && !this.isLoadingData) {
        this.isLoadingData = true;
        const url = editor.getActionUrl(editor.getOpt('imageManagerActionName'));
          const isJsonp = false;
        ajax.request(url, {
          'timeout': 100_000,
          'dataType': isJsonp ? 'jsonp' : '',
          'headers': editor.options.serverHeaders || {},
          'data': utils.extend({
            start: this.listIndex,
            size: this.listSize
          }, editor.queryCommandValue('serverparam')),
          'method': 'get',
          'onsuccess': function (r) {
            try {
              let json = isJsonp ? r : eval(`(${  r.responseText  })`);
              json = editor.options.serverResponsePrepare(json);
              if (json.state === 'SUCCESS') {
                _this.pushData(json.list);
                _this.listIndex = Number.parseInt(json.start) + Number.parseInt(json.list.length);
                if (_this.listIndex >= json.total) {
                  _this.listEnd = true;
                }
                _this.isLoadingData = false;
              }
            } catch {
              if (r.responseText.includes('ue_separate_ue')) {
                const list = r.responseText.split(r.responseText);
                _this.pushData(list);
                _this.listIndex = Number.parseInt(list.length);
                _this.listEnd = true;
                _this.isLoadingData = false;
              }
            }
          },
          'onerror': function () {
            _this.isLoadingData = false;
          }
        });
      }
    },
    /* 添加图片到列表界面上 */
    pushData (list) {
      const _this = this; let i; let icon; let img; let item;
        const urlPrefix = editor.getOpt('imageManagerUrlPrefix');
      for (i = 0; i < list.length; i++) {
        if (list[i] && list[i].url) {
          item = document.createElement('li');
          img = document.createElement('img');
          icon = document.createElement('span');

          domUtils.on(img, 'load', (function (image) {
            return function () {
              _this.scale(image, image.parentNode.offsetWidth, image.parentNode.offsetHeight);
            }
          })(img));
          img.width = 113;
          img.setAttribute('src', urlPrefix + list[i].url + (list[i].url.includes('?') ? '&noCache=' : '?noCache=') + (Date.now()).toString(36));
          img.setAttribute('_src', urlPrefix + list[i].url);
          domUtils.addClass(icon, 'icon');

          item.append(img);
          item.append(icon);
          this.list.insertBefore(item, this.clearFloat);
        }
      }
    },
    /* 改变图片大小 */
    scale (img, w, h, type) {
      const oh = img.height;
        const ow = img.width;

      if (type == 'justify') {
        if (ow >= oh) {
          img.width = w;
          img.height = h * oh / ow;
          img.style.marginLeft = `-${  Number.parseInt((img.width - w) / 2)  }px`;
        } else {
          img.width = w * ow / oh;
          img.height = h;
          img.style.marginTop = `-${  Number.parseInt((img.height - h) / 2)  }px`;
        }
      } else {
        if (ow >= oh) {
          img.width = w * ow / oh;
          img.height = h;
          img.style.marginLeft = `-${  Number.parseInt((img.width - w) / 2)  }px`;
        } else {
          img.width = w;
          img.height = h * oh / ow;
          img.style.marginTop = `-${  Number.parseInt((img.height - h) / 2)  }px`;
        }
      }
    },
    getInsertList () {
      const align = getAlign(); let i; const lis = this.list.children; const list = [];
      for (i = 0; i < lis.length; i++) {
        if (domUtils.hasClass(lis[i], 'selected')) {
          const img = lis[i].firstChild;
            const src = img.getAttribute('_src');
          list.push({
            src,
            _src: src,
            alt: src.slice(src.lastIndexOf('/') + 1),
            floatStyle: align
          });
        }

      }
      return list;
    }
  };

})();
