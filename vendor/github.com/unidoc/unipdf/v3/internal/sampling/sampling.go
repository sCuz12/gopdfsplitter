//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package sampling ;import (_g "github.com/unidoc/unipdf/v3/internal/bitwise";_f "github.com/unidoc/unipdf/v3/internal/imageutil";_c "io";);func (_bgf *Writer )WriteSamples (samples []uint32 )error {for _dcg :=0;_dcg < len (samples );_dcg ++{if _ag :=_bgf .WriteSample (samples [_dcg ]);
_ag !=nil {return _ag ;};};return nil ;};type SampleWriter interface{WriteSample (_ba uint32 )error ;WriteSamples (_bf []uint32 )error ;};type Writer struct{_af _f .ImageBase ;_gb *_g .Writer ;_cga ,_dfd int ;_add bool ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _eea []uint32 ;
_cg :=bitsPerOutputSample ;var _gee uint32 ;var _ac uint32 ;_cbfb :=0;_ad :=0;_dc :=0;for _dc < len (data ){if _cbfb > 0{_beb :=_cbfb ;if _cg < _beb {_beb =_cg ;};_gee =(_gee <<uint (_beb ))|(_ac >>uint (bitsPerInputSample -_beb ));_cbfb -=_beb ;if _cbfb > 0{_ac =_ac <<uint (_beb );
}else {_ac =0;};_cg -=_beb ;if _cg ==0{_eea =append (_eea ,_gee );_cg =bitsPerOutputSample ;_gee =0;_ad ++;};}else {_ff :=data [_dc ];_dc ++;_ca :=bitsPerInputSample ;if _cg < _ca {_ca =_cg ;};_cbfb =bitsPerInputSample -_ca ;_gee =(_gee <<uint (_ca ))|(_ff >>uint (_cbfb ));
if _ca < bitsPerInputSample {_ac =_ff <<uint (_ca );};_cg -=_ca ;if _cg ==0{_eea =append (_eea ,_gee );_cg =bitsPerOutputSample ;_gee =0;_ad ++;};};};for _cbfb >=bitsPerOutputSample {_db :=_cbfb ;if _cg < _db {_db =_cg ;};_gee =(_gee <<uint (_db ))|(_ac >>uint (bitsPerInputSample -_db ));
_cbfb -=_db ;if _cbfb > 0{_ac =_ac <<uint (_db );}else {_ac =0;};_cg -=_db ;if _cg ==0{_eea =append (_eea ,_gee );_cg =bitsPerOutputSample ;_gee =0;_ad ++;};};if _cg > 0&&_cg < bitsPerOutputSample {_gee <<=uint (_cg );_eea =append (_eea ,_gee );};return _eea ;
};func NewWriter (img _f .ImageBase )*Writer {return &Writer {_gb :_g .NewWriterMSB (img .Data ),_af :img ,_dfd :img .ColorComponents ,_add :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};type Reader struct{_cb _f .ImageBase ;
_cc *_g .Reader ;_e ,_d ,_geg int ;_da bool ;};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_ge []uint32 )error ;};func (_dab *Writer )WriteSample (sample uint32 )error {if _ ,_eeb :=_dab ._gb .WriteBits (uint64 (sample ),_dab ._af .BitsPerComponent );
_eeb !=nil {return _eeb ;};_dab ._dfd --;if _dab ._dfd ==0{_dab ._dfd =_dab ._af .ColorComponents ;_dab ._cga ++;};if _dab ._cga ==_dab ._af .Width {if _dab ._add {_dab ._gb .FinishByte ();};_dab ._cga =0;};return nil ;};func (_ded *Reader )ReadSamples (samples []uint32 )(_a error ){for _ce :=0;
_ce < len (samples );_ce ++{samples [_ce ],_a =_ded .ReadSample ();if _a !=nil {return _a ;};};return nil ;};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _fc []uint32 ;_ee :=bitsPerSample ;var _cbf uint32 ;var _gf byte ;_df :=0;_cf :=0;
_cfd :=0;for _cfd < len (data ){if _df > 0{_bc :=_df ;if _ee < _bc {_bc =_ee ;};_cbf =(_cbf <<uint (_bc ))|uint32 (_gf >>uint (8-_bc ));_df -=_bc ;if _df > 0{_gf =_gf <<uint (_bc );}else {_gf =0;};_ee -=_bc ;if _ee ==0{_fc =append (_fc ,_cbf );_ee =bitsPerSample ;
_cbf =0;_cf ++;};}else {_cbb :=data [_cfd ];_cfd ++;_daa :=8;if _ee < _daa {_daa =_ee ;};_df =8-_daa ;_cbf =(_cbf <<uint (_daa ))|uint32 (_cbb >>uint (_df ));if _daa < 8{_gf =_cbb <<uint (_daa );};_ee -=_daa ;if _ee ==0{_fc =append (_fc ,_cbf );_ee =bitsPerSample ;
_cbf =0;_cf ++;};};};for _df >=bitsPerSample {_aa :=_df ;if _ee < _aa {_aa =_ee ;};_cbf =(_cbf <<uint (_aa ))|uint32 (_gf >>uint (8-_aa ));_df -=_aa ;if _df > 0{_gf =_gf <<uint (_aa );}else {_gf =0;};_ee -=_aa ;if _ee ==0{_fc =append (_fc ,_cbf );_ee =bitsPerSample ;
_cbf =0;_cf ++;};};return _fc ;};func NewReader (img _f .ImageBase )*Reader {return &Reader {_cc :_g .NewReader (img .Data ),_cb :img ,_geg :img .ColorComponents ,_da :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};func (_bg *Reader )ReadSample ()(uint32 ,error ){if _bg ._d ==_bg ._cb .Height {return 0,_c .EOF ;
};_de ,_be :=_bg ._cc .ReadBits (byte (_bg ._cb .BitsPerComponent ));if _be !=nil {return 0,_be ;};_bg ._geg --;if _bg ._geg ==0{_bg ._geg =_bg ._cb .ColorComponents ;_bg ._e ++;};if _bg ._e ==_bg ._cb .Width {if _bg ._da {_bg ._cc .ConsumeRemainingBits ();
};_bg ._e =0;_bg ._d ++;};return uint32 (_de ),nil ;};